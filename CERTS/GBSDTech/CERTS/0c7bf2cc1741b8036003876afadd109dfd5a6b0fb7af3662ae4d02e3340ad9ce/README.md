# STIR/SHAKEN CA Ecosystem Compliance

## Certificate MYPBXManager SHAKEN

Tested At: 17 Dec 22 00:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 194 day(s)\
Subject: CN=MYPBXManager SHAKEN, O=MYPBXManager LLC, ST=New York, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://www.mypbxmanager.net/local/PubShakenCert.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIClTCCAjqgAwIBAgICEO4wCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMjA2MjgxODEzNDlaFw0yMzA2MjgxODEzNDlaMFkxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEZMBcGA1UECgwQTVlQQlhNYW5hZ2VyIExMQzEcMBoGA1UEAwwTTVlQQlhNYW5hZ2VyIFNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOiNN5L%2FJJEZJyD%2FZOPByvanFEICg5i3zzo%2F%2B61M3tZuNJ2ZwfiSlgnmaEg90%2FbYwJaIjv7vLvx%2F39Arg163jVSjgeowgecwFQYIKwYBBQUHARoECTAIoAYWNTU4SjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBTEMewZvvxA%2FkxsQZf7hFbSvEBsIjCBkAYDVR0jBIGIMIGFgBRFCSVsXS7H4VCeIIYTuSTGc0%2FiPqFppGcwZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMREwDwYDVQQHDAhGdCBXb3J0aDERMA8GA1UECgwIR0JTRFRlY2gxIDAeBgNVBAMMF0dCU0RUZWNoIFNIQUtFTiBSb290IENBggIQADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAIto2IChvR5g7bKhGR%2Fw%2B1luDuQMPNNnpxlzBafAtPLnAiEAsYVsntQTEffzS3sQoj2tlNUwBu90ooGphR46DK2pX98%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_serial_number](../../ISSUES/e_atis_serial_number/README.md) | error | ATIS1000080 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |


Generated: 17 Dec 22 00:12 UTC