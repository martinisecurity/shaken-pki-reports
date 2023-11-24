# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 24 Nov 23 11:01 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -431 day(s)\
Subject: CN=SHAKEN 5512, OU=SHAKEN, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/617f8185f2c473295280af2549b105ba.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqKgAwIBAgIQYuKYHmLh6xfCVp%2FpdkVs5zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTEyMDE1MjRaFw0yMjA5MTgyMDE1MjNaMFkxCzAJBgNVBAYTAlVTMSMwIQYDVQQKExpBbmRyZXcgV2FyZCBDb25zdWx0aW5nIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPQvmZWTEYJBRc9ZhThEFP1sC7%2BQAtTXZyqSGpBOhALP%2Bwtaa3k%2BVzxrF5XzRQyzBD791%2BL4bOJ6Z4VzcEQxtJejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUsHHKtKppNGuIcdvev%2FZbHH8DVBQwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxMjAKBggqhkjOPQQDAgNHADBEAiBWR7oCTdz5tmDrKX4NL5QK3pElfgDZQz18wwtRpfssKwIgdMRajs4jsiJtsFGaOIn9P6%2Bfrkv0521gC3FCY8xWzEs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC