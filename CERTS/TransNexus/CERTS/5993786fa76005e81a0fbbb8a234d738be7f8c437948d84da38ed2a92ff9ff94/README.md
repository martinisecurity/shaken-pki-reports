# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 26 Aug 24 17:45 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -557 day(s)\
Subject: CN=SHAKEN 738J, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/c559cb31-c1c4-480c-acde-f1bc4cfbb811.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIQWkcTWBuAQQ4%2BMSXFPcCRFTAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjA2MjAzNjA0WhcNMjMwMjE1MjAzNjAzWjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJU2t5U3dpdGNoMRQwEgYDVQQDEwtTSEFLRU4gNzM4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMm0ZGIN9nKCQGiUgflVcPNcqAwvDXxASvYywQcjQEqams73sOLkd6RDi7xc3L7La84XreE52beMdt7flNT2rQajggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUoaEgz8NCe6eZ3VjascYrQnAiQ7MwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzM4SjAKBggqhkjOPQQDAgNIADBFAiB7Y8ddatePTtPsHGOygD235556q4UT4Mo5OrxOui7kOQIhAI7KtUBrl3nOdeOwgovPdfSTDIrLBNn5B3c1dKISneyk)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC