# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 8526

Tested At: 21 Nov 23 16:46 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -321 day(s)\
Subject: CN=SHAKEN 8526, O=MetTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/8526/65708bb9-ff7b-412e-a823-edf04c433747.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICxjCCAmygAwIBAgIQQbY8keeeaiNlbuPzY2jYJjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjA1MTUyMzIwWhcNMjMwMTA0MTUyMzE5WjA0MQswCQYDVQQGEwJVUzEPMA0GA1UEChMGTWV0VGVsMRQwEgYDVQQDEwtTSEFLRU4gODUyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA8%2BxOssPifg0SlFbg4wm5aZvAFJUQNvfOtdcJtRVxMrsOKbLbXHw3%2FtYSOeuBRcbyVzE7KLu6es0NsRi0nBjlGjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUSKv713SNf5HrhllyoPo5p1fcboYwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEODUyNjAKBggqhkjOPQQDAgNIADBFAiEAi0VCWb1Xqknp2e541rYih4TtRlQkaes%2FBtgWjwyUVFICIFGYHOnKsmrLvoOEFKCHnpFA7P9zh8b2PhB6FGIJw9qA)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC