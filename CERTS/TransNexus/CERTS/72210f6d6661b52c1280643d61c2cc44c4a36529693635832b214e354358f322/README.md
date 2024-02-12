# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 056K

Tested At: 12 Feb 24 18:55 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -476 day(s)\
Subject: CN=SHAKEN 056K, OU=SHAKEN, O=Logista Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d673a99f-ad2d-4256-8d65-e20ff91adba4/8140550d516aca64f042df75c37894ed.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQab3wmvHEqE8nJHHHPJTrlTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcxODM1MjZaFw0yMjEwMjQxODM1MjVaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFMb2dpc3RhIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDU2SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNLF0Q%2FXpafRl%2Bs7VBR7P0MyrxN7IEgHoVAgy5tUbVXDmpt8EsKzbL7KE6EGltke6EIHvdeThWDuOhUnyz%2F4Ve6jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUSBGTAHOkwPXFllzArjuuUeYoFKMwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDU2SzAKBggqhkjOPQQDAgNHADBEAiBiHzgKXRThScvXnFHQFxKUPpKf891jFzZpFIp%2B79YcIgIgErQVfeJSQK3%2FDWpflOdvp8i5o3mdp5t6Ox8r9cSgzuo%3D)

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


Generated: 12 Feb 24 19:26 UTC