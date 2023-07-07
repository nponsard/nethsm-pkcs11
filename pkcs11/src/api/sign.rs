use cryptoki_sys::CKR_OK;
use log::{error, trace};

use crate::{
    backend::mechanism::{CkRawMechanism, Mechanism},
    data::SESSION_MANAGER,
    lock_mutex,
};

pub extern "C" fn C_SignInit(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pMechanism: *mut cryptoki_sys::CK_MECHANISM,
    hKey: cryptoki_sys::CK_OBJECT_HANDLE,
) -> cryptoki_sys::CK_RV {
    trace!("C_SignInit() called with hKey {} and mech", hKey);
    if pMechanism.is_null() {
        return cryptoki_sys::CKR_ARGUMENTS_BAD;
    }
    trace!("C_SignInit() mech: {:?}", unsafe { *pMechanism });

    let raw_mech = unsafe { CkRawMechanism::from_raw_ptr_unchecked(pMechanism) };

    let mech = match Mechanism::from_ckraw_mech(&raw_mech) {
        Ok(mech) => mech,
        Err(e) => {
            error!("C_SignInit() failed to convert mechanism: {:?}", e);
            return cryptoki_sys::CKR_MECHANISM_INVALID;
        }
    };

    let mut manager = lock_mutex!(SESSION_MANAGER);

    let session = match manager.get_session_mut(hSession) {
        Some(session) => session,
        None => {
            error!(
                "C_SignInit() called with invalid session handle {}.",
                hSession
            );
            return cryptoki_sys::CKR_SESSION_HANDLE_INVALID;
        }
    };

    session.sign_init(&mech, hKey)
}

pub extern "C" fn C_Sign(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pData: *mut cryptoki_sys::CK_BYTE,
    ulDataLen: cryptoki_sys::CK_ULONG,
    pSignature: *mut cryptoki_sys::CK_BYTE,
    pulSignatureLen: *mut cryptoki_sys::CK_ULONG,
) -> cryptoki_sys::CK_RV {
    trace!("C_Sign() called");

    let mut manager = lock_mutex!(SESSION_MANAGER);

    let session = match manager.get_session_mut(hSession) {
        Some(session) => session,
        None => {
            error!("C_Sign() called with invalid session handle {}.", hSession);
            return cryptoki_sys::CKR_SESSION_HANDLE_INVALID;
        }
    };

    // TODO: if pSignature is NULL, then this function should return the maximum length of the signature
    if pData.is_null() || pSignature.is_null() || pulSignatureLen.is_null() {
        session.sign_clear();
        return cryptoki_sys::CKR_ARGUMENTS_BAD;
    }

    let data = unsafe { std::slice::from_raw_parts(pData, ulDataLen as usize) };

    let signature = match session.sign(data) {
        Ok(signature) => signature,
        Err(err) => return err,
    };

    if signature.len() > unsafe { *pulSignatureLen } as usize {
        unsafe {
            std::ptr::write(pulSignatureLen, signature.len() as u64);
        }
        return cryptoki_sys::CKR_BUFFER_TOO_SMALL;
    }

    unsafe {
        std::ptr::copy_nonoverlapping(signature.as_ptr(), pSignature, signature.len());
        std::ptr::write(pulSignatureLen, signature.len() as u64);
    }

    session.sign_clear();

    cryptoki_sys::CKR_OK
}

pub extern "C" fn C_SignUpdate(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pPart: *mut cryptoki_sys::CK_BYTE,
    ulPartLen: cryptoki_sys::CK_ULONG,
) -> cryptoki_sys::CK_RV {
    trace!("C_SignUpdate() called");

    let mut manager = lock_mutex!(SESSION_MANAGER);

    let session = match manager.get_session_mut(hSession) {
        Some(session) => session,
        None => {
            error!("C_Sign() called with invalid session handle {}.", hSession);
            return cryptoki_sys::CKR_SESSION_HANDLE_INVALID;
        }
    };

    if pPart.is_null() {
        session.sign_clear();
        return cryptoki_sys::CKR_ARGUMENTS_BAD;
    }

    let part = unsafe { std::slice::from_raw_parts(pPart, ulPartLen as usize) };

    match session.sign_update(part) {
        Ok(_) => cryptoki_sys::CKR_OK,
        Err(err) => {
            session.sign_clear();
            err
        }
    }
}

pub extern "C" fn C_SignFinal(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pSignature: *mut cryptoki_sys::CK_BYTE,
    pulSignatureLen: *mut cryptoki_sys::CK_ULONG,
) -> cryptoki_sys::CK_RV {
    trace!("C_SignFinal() called");

    let mut manager = lock_mutex!(SESSION_MANAGER);

    let session = match manager.get_session_mut(hSession) {
        Some(session) => session,
        None => {
            error!("C_Sign() called with invalid session handle {}.", hSession);
            return cryptoki_sys::CKR_SESSION_HANDLE_INVALID;
        }
    };

    // TODO: if pSignature is NULL, then this function should return the maximum length of the signature
    if pSignature.is_null() || pulSignatureLen.is_null() {
        session.sign_clear();
        return cryptoki_sys::CKR_ARGUMENTS_BAD;
    }

    let signature = match session.sign_final() {
        Ok(signature) => signature,
        Err(err) => {
            session.sign_clear();
            return err;
        }
    };

    unsafe {
        std::ptr::copy_nonoverlapping(signature.as_ptr(), pSignature, signature.len());
        std::ptr::write(pulSignatureLen, signature.len() as u64);
    }
    session.sign_clear();

    CKR_OK
}

pub extern "C" fn C_SignRecoverInit(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pMechanism: cryptoki_sys::CK_MECHANISM_PTR,
    hKey: cryptoki_sys::CK_OBJECT_HANDLE,
) -> cryptoki_sys::CK_RV {
    trace!("C_SignRecoverInit() called");
    cryptoki_sys::CKR_FUNCTION_NOT_SUPPORTED
}

pub extern "C" fn C_SignRecover(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pData: cryptoki_sys::CK_BYTE_PTR,
    ulDataLen: cryptoki_sys::CK_ULONG,
    pSignature: cryptoki_sys::CK_BYTE_PTR,
    pulSignatureLen: cryptoki_sys::CK_ULONG_PTR,
) -> cryptoki_sys::CK_RV {
    trace!("C_SignRecover() called");
    cryptoki_sys::CKR_FUNCTION_NOT_SUPPORTED
}

pub extern "C" fn C_SignEncryptUpdate(
    hSession: cryptoki_sys::CK_SESSION_HANDLE,
    pPart: cryptoki_sys::CK_BYTE_PTR,
    ulPartLen: cryptoki_sys::CK_ULONG,
    pEncryptedPart: cryptoki_sys::CK_BYTE_PTR,
    pulEncryptedPartLen: cryptoki_sys::CK_ULONG_PTR,
) -> cryptoki_sys::CK_RV {
    trace!("C_SignEncryptUpdate() called");
    cryptoki_sys::CKR_FUNCTION_NOT_SUPPORTED
}