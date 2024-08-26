# STIR/SHAKEN CA Ecosystem Compliance

## Certificate STIR ECC PNCB

Tested At: 26 Aug 24 18:14 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 268 day(s)\
Subject: CN=STIR ECC PNCB, O=PNC Bank, C=US\
Issuer: CN=Neustar UAT Enterprise Certified Caller Intermediate CA1, OU=www.ccid-uat.neustar, O=Neustar Inc a TransUnion company, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11617.10439.pem

[View certificate details](https://x509.io/?cert=MIIC0jCCAnigAwIBAgIUd9HfDxnOn3sdz3prYrC8CnsUhokwCgYIKoZIzj0EAwIwgZoxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluYyBhIFRyYW5zVW5pb24gY29tcGFueTEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxQTA%2FBgNVBAMMOE5ldXN0YXIgVUFUIEVudGVycHJpc2UgQ2VydGlmaWVkIENhbGxlciBJbnRlcm1lZGlhdGUgQ0ExMB4XDTI0MDUyMTE0Mzg1NVoXDTI1MDUyMTE0Mzg1NVowODELMAkGA1UEBhMCVVMxETAPBgNVBAoMCFBOQyBCYW5rMRYwFAYDVQQDDA1TVElSIEVDQyBQTkNCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMY%2BwvT0%2F0C3Y7XZ1oy923Kklh6De2qBDh6pFhH%2BVuUqK%2BAJY%2FOsANOR7atmXxhNA%2FhqrEdjIg6oi329YlMWDuKOB%2FDCB%2BTAWBggrBgEFBQcBGgQKMAigBhYEUE5DQjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFCt9OLm%2Fe3nm%2F9WmsNzYwh%2FjlrGDMBUGA1UdIAQOMAwwCgYIKwYBBAHqAGQwagYDVR0fBGMwYTBfoF2gW4ZZaHR0cHM6Ly9jcmwtdWF0LWVjY2lkLmNjaWQubmV1c3Rhci5iaXovTmV1c3RhckVudGVycHJpc2VDZXJ0aWZpZWRDYWxsZXJJZEludGVybWVkaWF0ZS5jcmwwHQYDVR0OBBYEFDyIDyR9p39sBYaKtxEFKjui7ai5MA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiB5fyKJlengqkDQHl%2F9slj5sljPER7%2FTpZZhR%2Bq56fNQwIhALx8LfuP3mQiZRafAc2aT2EKd1%2FC1UkKFyU52xd605Qa)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN PNCB', but common name is 'STIR ECC PNCB' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.13568.100. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'STIR ECC PNCB' does not contain 'SHAKEN' |


Generated: 26 Aug 24 18:49 UTC