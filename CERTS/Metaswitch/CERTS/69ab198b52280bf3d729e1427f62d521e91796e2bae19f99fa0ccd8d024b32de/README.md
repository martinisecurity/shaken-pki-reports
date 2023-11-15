# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Pembroke Telephone Company, Inc SHAKEN Cert 0376

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 641 day(s)\
Subject: CN=Pembroke Telephone Company\\, Inc SHAKEN Cert 0376, O=Pembroke Telephone Company\\, Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/54f1b46e13a9f05d1b3f49b6f71314b0b97cc086

[View certificate details](https://understandingwebpki.com/?cert=MIICeDCCAh%2BgAwIBAgIQA874mQvHIqFVhDb1nSnekzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDgxODEwMzkxOVoXDTI1MDgxNzEwMzkxOVowcjELMAkGA1UEBhMCVVMxKDAmBgNVBAoMH1BlbWJyb2tlIFRlbGVwaG9uZSBDb21wYW55LCBJbmMxOTA3BgNVBAMMMFBlbWJyb2tlIFRlbGVwaG9uZSBDb21wYW55LCBJbmMgU0hBS0VOIENlcnQgMDM3NjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDRXIWMU01phWBCppA232ss8%2FvK6fWcyomgrmMzz%2FJcDTybXDvNrfr%2BHTc5l7DZ%2BQ6u2UOTQXSeX6rinVTJOA%2FajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDAzNzYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU5h4Hc4h6OLtwSJa4Of4pqTs9M18wHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgajNjehj%2FpsCD%2FuETi0%2B7xXNuonKopJOxoXY7EPhIY%2F0CIAed4NdBK5GBbWclg%2FitlAAk8fbM6rEGf7OlxK4IUw8v)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0376' |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


Generated: 15 Nov 23 17:17 UTC