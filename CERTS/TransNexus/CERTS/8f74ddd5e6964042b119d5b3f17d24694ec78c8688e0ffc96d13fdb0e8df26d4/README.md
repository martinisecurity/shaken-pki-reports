# STIR/SHAKEN CA Ecosystem Compliance

## Certificate MobileSphere SHAKEN 873J

Tested At: 15 Nov 23 18:01 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -306 day(s)\
Subject: CN=MobileSphere SHAKEN 873J, OU=MobileSphere, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/873J/ac03ff61-e0df-456a-b2b1-d4a8d57afcff.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICnTCCAkOgAwIBAgIQeiz1myI12aKdpmKzEW8RjTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAxMTMxNDE3NDlaFw0yMzAxMTMxNDE3NDhaMFwxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRUwEwYDVQQLEwxNb2JpbGVTcGhlcmUxITAfBgNVBAMTGE1vYmlsZVNwaGVyZSBTSEFLRU4gODczSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSxTI2Ux3uuYpwmsOAzhvXJWEGbFY%2F2GZpzBMUIKc4liYR2WB05DMHDsuPLXgEo3XdNAC1GYW58ZHWvUAuOYWqjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAoMPwyCwaCLlPdDsY%2F4d89CKWXlMB8GA1UdIwQYMBaAFJSGOT%2Fk9ZKgn%2F%2FcJ3UamkbweeFiMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFgYIKwYBBQUHARoECjAIoAYWBDg3M0owCgYIKoZIzj0EAwIDSAAwRQIhAJz4MfP%2FCxs8hqNh6LI3q8c0PkUOMoNRWrZURuwRMQH3AiBNUMO4%2B9xP0guAli2%2BqNKLD0W7zWJ2Zf2LxJ%2BEjLdMmQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 15 Nov 23 18:10 UTC