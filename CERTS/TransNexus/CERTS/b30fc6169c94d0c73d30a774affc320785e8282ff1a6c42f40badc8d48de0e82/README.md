# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 733J

Tested At: 22 Aug 24 15:19 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -559 day(s)\
Subject: CN=SHAKEN 733J, O=Wiretap Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/257dc117-4652-4444-9058-fceb2d81e92c/179f4198b823db282c0499fa6bd45637.pem

[View certificate details](https://x509.io/?cert=MIICzzCCAnWgAwIBAgIQQu06D7qpx5FC0YwOAk40RjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjAyMjAyODM3WhcNMjMwMjA5MjAyODM2WjA9MQswCQYDVQQGEwJVUzEYMBYGA1UEChMPV2lyZXRhcCBUZWxlY29tMRQwEgYDVQQDEwtTSEFLRU4gNzMzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLotsnGiFVr%2BiHZVfWRuDX%2BXFztUoANMkMipMm5FvzFvyEMbAyj1%2F7OCk%2FgHrhAgY6EyGhgJe8HVdmQYMFuCMd2jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUyuyENBcn4rFCO5%2BJYJP0MPJ91iUwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzMzSjAKBggqhkjOPQQDAgNIADBFAiEAkOe0AT0LOVHn3NqkIdtuyq9T4ka7KyLIngs%2FZ3uDfvgCIAzBHhuiowMGWyxf55FbomZun4hFZ5sr6IyP7TTpSH0M)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 16:06 UTC