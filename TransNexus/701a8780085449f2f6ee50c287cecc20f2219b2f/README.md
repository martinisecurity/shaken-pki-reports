# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 701a8780085449f2f6ee50c287cecc20f2219b2f
Tested At: 2022-10-26 20:21:35 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 79 day(s)\
Subject: CN=MobileSphere SHAKEN 873J, OU=MobileSphere, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/873J/ac03ff61-e0df-456a-b2b1-d4a8d57afcff.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICnTCCAkOgAwIBAgIQeiz1myI12aKdpmKzEW8RjTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAxMTMxNDE3NDlaFw0yMzAxMTMxNDE3NDhaMFwxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRUwEwYDVQQLEwxNb2JpbGVTcGhlcmUxITAfBgNVBAMTGE1vYmlsZVNwaGVyZSBTSEFLRU4gODczSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSxTI2Ux3uuYpwmsOAzhvXJWEGbFY%2F2GZpzBMUIKc4liYR2WB05DMHDsuPLXgEo3XdNAC1GYW58ZHWvUAuOYWqjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAoMPwyCwaCLlPdDsY%2F4d89CKWXlMB8GA1UdIwQYMBaAFJSGOT%2Fk9ZKgn%2F%2FcJ3UamkbweeFiMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFgYIKwYBBQUHARoECjAIoAYWBDg3M0owCgYIKoZIzj0EAwIDSAAwRQIhAJz4MfP%2FCxs8hqNh6LI3q8c0PkUOMoNRWrZURuwRMQH3AiBNUMO4%2B9xP0guAli2%2BqNKLD0W7zWJ2Zf2LxJ%2BEjLdMmQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_sti_issuer_dn
- e_sti_subject
- e_sti_crl_distribution
- e_sti_certificate_policies
- e_sti_extension_unknown
- e_sti_subject_cn
- n_sti_certificate_policy_critical
- e_sti_signature_algorithm
- e_sti_serial_number
- e_sti_key_usage
- e_sti_subject_public_key
- e_sti_authority_key_identifier
- e_sti_tn_auth_list
- e_sti_version
- e_sti_subject_key_identifier

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 26/10/2022 at 20:22:11