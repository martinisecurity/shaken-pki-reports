# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 733J

Tested At: 04 Oct 24 15:34 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -605 day(s)\
Subject: CN=SHAKEN 733J, O=Wiretap Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/257dc117-4652-4444-9058-fceb2d81e92c/1fe40767ac46886758647a6ac3e49092.pem

[View certificate details](https://x509.io/?cert=MIICzzCCAnWgAwIBAgIQW4Q31feJO9byIGbt8pvzjzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMTMwMjAyODM0WhcNMjMwMjA2MjAyODMzWjA9MQswCQYDVQQGEwJVUzEYMBYGA1UEChMPV2lyZXRhcCBUZWxlY29tMRQwEgYDVQQDEwtTSEFLRU4gNzMzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM%2B%2BIfKUOh%2F2fpAPDsW%2BUc%2FcN7dTq5wQ7AY89JhglbNHbgJDRhuxS%2BrYvZ4QXcvTneG1aX%2BaVKDKCU8F082jYG6jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUcXRZGm9PQOmC%2BFVx7jg0HuRpzdAwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzMzSjAKBggqhkjOPQQDAgNIADBFAiEA141e%2FHOmAOsN90VYBMzFgCP%2FiLVl7FdOrNmCoZZMBaQCIC%2F5A7FdLzEAVcVZv4jS%2FaPoU9QatFEUN32IfnHDdL9B)

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