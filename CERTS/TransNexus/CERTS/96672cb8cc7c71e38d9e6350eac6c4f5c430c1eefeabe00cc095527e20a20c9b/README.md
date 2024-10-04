# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 8526

Tested At: 04 Oct 24 15:50 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -278 day(s)\
Subject: CN=SHAKEN 8526, O=MetTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/8526/be6281a4-be9c-47d8-98fb-4371d09ddd53.pem

[View certificate details](https://x509.io/?cert=MIICxzCCAmygAwIBAgIQZxkBqhOLuS8%2F9S%2BoG9Ch%2BjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjMxMDMzMzE4WhcNMjMxMjMxMDMzMzE3WjA0MQswCQYDVQQGEwJVUzEPMA0GA1UEChMGTWV0VGVsMRQwEgYDVQQDEwtTSEFLRU4gODUyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIUnwWc%2FTSduW%2B5mBLDOa%2BaFuqaNvABA0VCyXmWtig%2BaDBn51DmJi1H%2BBzDrZzydBafNnPRG8d9Kamev06S%2FKqWjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUOhVnOCKHOATMZnIblECAPnKd3KQwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEODUyNjAKBggqhkjOPQQDAgNJADBGAiEAraAZ3cTb%2FnQpgFdpimMjvCUP0JumleEJbb97%2FA9lP38CIQC2Ih5KNQ%2FD1cIdtibyqmdZq%2BwW8%2FN9CVSviNwP7LSUVA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC