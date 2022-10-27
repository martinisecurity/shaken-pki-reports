# STIR/SHAKEN CA Ecosystem Compliance
## GBSDTech

### Certificate a3872afd09406d2745d204893b6b52bbf6380f84
Tested At: 2022-10-27 18:24:46 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 244 day(s)\
Subject: CN=MYPBXManager SHAKEN, O=MYPBXManager LLC, ST=New York, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US

Link: https://www.mypbxmanager.net/local/PubShakenCert.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIClTCCAjqgAwIBAgICEO4wCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMjA2MjgxODEzNDlaFw0yMzA2MjgxODEzNDlaMFkxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEZMBcGA1UECgwQTVlQQlhNYW5hZ2VyIExMQzEcMBoGA1UEAwwTTVlQQlhNYW5hZ2VyIFNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOiNN5L%2FJJEZJyD%2FZOPByvanFEICg5i3zzo%2F%2B61M3tZuNJ2ZwfiSlgnmaEg90%2FbYwJaIjv7vLvx%2F39Arg163jVSjgeowgecwFQYIKwYBBQUHARoECTAIoAYWNTU4SjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBTEMewZvvxA%2FkxsQZf7hFbSvEBsIjCBkAYDVR0jBIGIMIGFgBRFCSVsXS7H4VCeIIYTuSTGc0%2FiPqFppGcwZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMREwDwYDVQQHDAhGdCBXb3J0aDERMA8GA1UECgwIR0JTRFRlY2gxIDAeBgNVBAMMF0dCU0RUZWNoIFNIQUtFTiBSb290IENBggIQADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAIto2IChvR5g7bKhGR%2Fw%2B1luDuQMPNNnpxlzBafAtPLnAiEAsYVsntQTEffzS3sQoj2tlNUwBu90ooGphR46DK2pX98%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_tn_auth_list | error | ATIS-1000080v4 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| e_sti_serial_number | error | ATIS-1000080v4 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| e_sti_subject_cn | error | ATIS-1000080v4 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |

Generated: 27/10/2022 at 18:24:52