/*
 * NetHSM
 *
 * All endpoints expect exactly the specified JSON. Additional properties will cause a Bad Request Error (400). All HTTP errors contain a JSON structure with an explanation of type string. All [base64](https://tools.ietf.org/html/rfc4648#section-4) encoded values are Big Endian. 
 *
 * The version of the OpenAPI document: v1
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct KeyPrivateData {
    #[serde(rename = "primeP", skip_serializing_if = "Option::is_none")]
    pub prime_p: Option<String>,
    #[serde(rename = "primeQ", skip_serializing_if = "Option::is_none")]
    pub prime_q: Option<String>,
    #[serde(rename = "publicExponent", skip_serializing_if = "Option::is_none")]
    pub public_exponent: Option<String>,
    #[serde(rename = "data", skip_serializing_if = "Option::is_none")]
    pub data: Option<String>,
}

impl KeyPrivateData {
    pub fn new() -> KeyPrivateData {
        KeyPrivateData {
            prime_p: None,
            prime_q: None,
            public_exponent: None,
            data: None,
        }
    }
}

