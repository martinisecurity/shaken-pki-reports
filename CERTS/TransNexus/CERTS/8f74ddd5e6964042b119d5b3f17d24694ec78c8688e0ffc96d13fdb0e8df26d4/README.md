# STIR/SHAKEN CA Ecosystem Compliance

## Certificate MobileSphere SHAKEN 873J

Tested At: 22 Aug 24 15:28 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -587 day(s)\
Subject: CN=MobileSphere SHAKEN 873J, OU=MobileSphere, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/873J/ac03ff61-e0df-456a-b2b1-d4a8d57afcff.pem

[View certificate details](https://x509.io/?cert=MIICnTCCAkOgAwIBAgIQeiz1myI12aKdpmKzEW8RjTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAxMTMxNDE3NDlaFw0yMzAxMTMxNDE3NDhaMFwxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRUwEwYDVQQLEwxNb2JpbGVTcGhlcmUxITAfBgNVBAMTGE1vYmlsZVNwaGVyZSBTSEFLRU4gODczSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSxTI2Ux3uuYpwmsOAzhvXJWEGbFY%2F2GZpzBMUIKc4liYR2WB05DMHDsuPLXgEo3XdNAC1GYW58ZHWvUAuOYWqjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAoMPwyCwaCLlPdDsY%2F4d89CKWXlMB8GA1UdIwQYMBaAFJSGOT%2Fk9ZKgn%2F%2FcJ3UamkbweeFiMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFgYIKwYBBQUHARoECjAIoAYWBDg3M0owCgYIKoZIzj0EAwIDSAAwRQIhAJz4MfP%2FCxs8hqNh6LI3q8c0PkUOMoNRWrZURuwRMQH3AiBNUMO4%2B9xP0guAli2%2BqNKLD0W7zWJ2Zf2LxJ%2BEjLdMmQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 16:06 UTC