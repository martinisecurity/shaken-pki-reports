# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Pembroke Telephone Company, Inc SHAKEN Cert 0376

Tested At: 21 Nov 23 01:22 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 636 day(s)\
Subject: CN=Pembroke Telephone Company\\, Inc SHAKEN Cert 0376, O=Pembroke Telephone Company\\, Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/54f1b46e13a9f05d1b3f49b6f71314b0b97cc086

[View certificate details](https://understandingwebpki.com/?cert=MIICeDCCAh%2BgAwIBAgIQA874mQvHIqFVhDb1nSnekzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDgxODEwMzkxOVoXDTI1MDgxNzEwMzkxOVowcjELMAkGA1UEBhMCVVMxKDAmBgNVBAoMH1BlbWJyb2tlIFRlbGVwaG9uZSBDb21wYW55LCBJbmMxOTA3BgNVBAMMMFBlbWJyb2tlIFRlbGVwaG9uZSBDb21wYW55LCBJbmMgU0hBS0VOIENlcnQgMDM3NjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDRXIWMU01phWBCppA232ss8%2FvK6fWcyomgrmMzz%2FJcDTybXDvNrfr%2BHTc5l7DZ%2BQ6u2UOTQXSeX6rinVTJOA%2FajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDAzNzYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU5h4Hc4h6OLtwSJa4Of4pqTs9M18wHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgajNjehj%2FpsCD%2FuETi0%2B7xXNuonKopJOxoXY7EPhIY%2F0CIAed4NdBK5GBbWclg%2FitlAAk8fbM6rEGf7OlxK4IUw8v)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0376' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC