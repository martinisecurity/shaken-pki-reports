# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 291K

Tested At: 04 Oct 24 15:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -718 day(s)\
Subject: CN=SHAKEN 291K, OU=SHAKEN, O=Hypercore Networks\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b66a7496-cdd4-46e3-b219-cb0ecdc80d22/f10a6e2814430717be9d18e253b564a3.pem

[View certificate details](https://x509.io/?cert=MIIC%2BTCCAp%2BgAwIBAgIQaJn0KABBNWoOqrtJP8ZULzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDkxNjMwMjVaFw0yMjEwMTYxNjMwMjRaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdIeXBlcmNvcmUgTmV0d29ya3MsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMjkxSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJg%2F30TWCS%2BqfSpPMWJobk68cGGr%2BAHMISYaBBIT74fKxSlV59LuolLDRLZZvZQlAXnnDi8%2FTV%2BIvfalO5%2BoVWyjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUfkcTMliJqcAc1pLjslPJmoLn9N4wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMjkxSzAKBggqhkjOPQQDAgNIADBFAiEAnWYo0OASTxHUyAtINNTow21y0b8uQarkR9GopzNzLbcCIHv%2BGg%2BNLZhzEn3x1emrMJ7kZpL0rIJ1FogQ7CLmVQbD)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC