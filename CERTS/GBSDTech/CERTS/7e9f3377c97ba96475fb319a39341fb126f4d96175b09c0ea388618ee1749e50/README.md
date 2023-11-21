# STIR/SHAKEN CA Ecosystem Compliance

## Certificate FracTEL LLC SHAKEN

Tested At: 21 Nov 23 01:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 218 day(s)\
Subject: CN=FracTEL LLC SHAKEN, O=FracTEL LLC, ST=Florida, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://ssc.fractel.net/ssc/fractelssc.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICjjCCAjOgAwIBAgICEPgwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMzA2MjYxODMxMjhaFw0yNDA2MjUxODMxMjhaMFIxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRQwEgYDVQQKDAtGcmFjVEVMIExMQzEbMBkGA1UEAwwSRnJhY1RFTCBMTEMgU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEcBfY53qnkli6xgx42KzkImD5rA%2FGKuw28EWX72DfwGMsO9a7NTShy60IvQSrP0zVVdMcfcyaWSiru%2Bvpy5ZJt6OB6jCB5zAVBggrBgEFBQcBGgQJMAigBhY5NjVIMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFCrRjwgHtJR6jdPADO7QoBnmQPAtMIGQBgNVHSMEgYgwgYWAFEUJJWxdLsfhUJ4ghhO5JMZzT%2BI%2BoWmkZzBlMQswCQYDVQQGEwJVUzEOMAwGA1UECAwFVGV4YXMxETAPBgNVBAcMCEZ0IFdvcnRoMREwDwYDVQQKDAhHQlNEVGVjaDEgMB4GA1UEAwwXR0JTRFRlY2ggU0hBS0VOIFJvb3QgQ0GCAhAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA010Z4M818FruAl7%2BRKisty4uBhGFBNqPpafT2wgI%2FBwCIQD8Yfl9tcU9u4yN1uxyZ2XoId%2BKc4oxDxBbYRxFsc74cg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | serial number size is less than 64 bits, got 16 bits |


Generated: 21 Nov 23 01:55 UTC