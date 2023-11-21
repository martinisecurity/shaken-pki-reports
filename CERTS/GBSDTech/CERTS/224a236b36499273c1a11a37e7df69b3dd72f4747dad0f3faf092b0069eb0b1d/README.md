# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Edify SHAKEN

Tested At: 21 Nov 23 01:51 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -201 day(s)\
Subject: CN=Edify SHAKEN, O=Edify, ST=California, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://sti.edifylabs.net/05-03-2022.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIChDCCAiqgAwIBAgICEMgwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMjA1MDMxMzQxMjhaFw0yMzA1MDMxMzQxMjhaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMQ4wDAYDVQQKDAVFZGlmeTEVMBMGA1UEAwwMRWRpZnkgU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqeiy9cf%2BQR1i2tL6fOOT9QKujt9Nvh0%2Be5lctanRu85lIpMrJcZeLEMOY%2BqLbAjXGoK1pNzOrrRo5U8M8lZDBqOB6jCB5zAVBggrBgEFBQcBGgQJMAigBhY4NjdKMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFDERhZb0KAPsp75iPkWcBE2E8VC3MIGQBgNVHSMEgYgwgYWAFEUJJWxdLsfhUJ4ghhO5JMZzT%2BI%2BoWmkZzBlMQswCQYDVQQGEwJVUzEOMAwGA1UECAwFVGV4YXMxETAPBgNVBAcMCEZ0IFdvcnRoMREwDwYDVQQKDAhHQlNEVGVjaDEgMB4GA1UEAwwXR0JTRFRlY2ggU0hBS0VOIFJvb3QgQ0GCAhAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiA4Lx6DsdQONvJP0WZXvDzY5FPds9jKEf76G7Ii%2BA3yOwIhANYp6bd%2B2aLwjXLxJnnp3pEMdNvo7tiEVCK9yuZYa0Kv)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC