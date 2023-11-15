# STIR/SHAKEN CA Ecosystem Compliance

## Certificate MYPBXManager SHAKEN

Tested At: 15 Nov 23 16:49 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 992 day(s)\
Subject: CN=MYPBXManager SHAKEN, O=MYPBXManager LLC, ST=New York, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://www.mypbxmanager.net/local/PubShakenCert.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIClDCCAjqgAwIBAgICEPwwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMzA4MDMxOTU1MzRaFw0yNjA4MDIxOTU1MzRaMFkxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEZMBcGA1UECgwQTVlQQlhNYW5hZ2VyIExMQzEcMBoGA1UEAwwTTVlQQlhNYW5hZ2VyIFNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKR5v9jj%2FXEwbdJeR1vA94RigXdQXlCDYqUjZXVvA6xD0yfYoSj%2Bxr6dapV%2BDm1ML4cZeyrvlmGQtSONAveI2oyjgeowgecwFQYIKwYBBQUHARoECTAIoAYWNTU4SjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRyXiwp14NaFyqSITK2I%2FPqbpnx%2BTCBkAYDVR0jBIGIMIGFgBRFCSVsXS7H4VCeIIYTuSTGc0%2FiPqFppGcwZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMREwDwYDVQQHDAhGdCBXb3J0aDERMA8GA1UECgwIR0JTRFRlY2gxIDAeBgNVBAMMF0dCU0RUZWNoIFNIQUtFTiBSb290IENBggIQADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgC2AmnBttVZr9Wq%2B1zCdOZHXwaj4JlPpkTisNnmBmXnMCIQCfh4cCaHkAH527wDaewbR871jYrKv3ycMYOMGDQJrBIg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_serial_number](../../ISSUES/e_atis_serial_number/README.md) | error | ATIS1000080 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |


Generated: 15 Nov 23 16:51 UTC