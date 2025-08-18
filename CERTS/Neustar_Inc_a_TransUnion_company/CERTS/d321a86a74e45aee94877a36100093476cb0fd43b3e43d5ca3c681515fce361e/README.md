# STIR/SHAKEN CA Ecosystem Compliance

## Certificate STIR ECC PNCB

Tested At: 18 Aug 25 20:20 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 261 day(s)\
Subject: CN=STIR ECC PNCB, O=PNC Bank, C=US\
Issuer: CN=Neustar UAT Enterprise Certified Caller Intermediate CA1, OU=www.ccid-uat.neustar, O=Neustar Inc a TransUnion company, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11617.10934.pem

[View certificate details](https://x509.io/?cert=MIIC0jCCAnigAwIBAgIUPD1cwcWlNFLQ0jKH%2FSSlyUDx8iQwCgYIKoZIzj0EAwIwgZoxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluYyBhIFRyYW5zVW5pb24gY29tcGFueTEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxQTA%2FBgNVBAMMOE5ldXN0YXIgVUFUIEVudGVycHJpc2UgQ2VydGlmaWVkIENhbGxlciBJbnRlcm1lZGlhdGUgQ0ExMB4XDTI1MDUwNjE1MDQyNVoXDTI2MDUwNjE1MDQyNVowODELMAkGA1UEBhMCVVMxETAPBgNVBAoMCFBOQyBCYW5rMRYwFAYDVQQDDA1TVElSIEVDQyBQTkNCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECsGBghk2TIyr5IOBCVIL8dVPShrLTR76%2FrEjTPfKTeYtq4ilocL%2B5HCcwlvW409IghCi7PJySf5EZpFTsYZ%2BTqOB%2FDCB%2BTAWBggrBgEFBQcBGgQKMAigBhYEUE5DQjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFCt9OLm%2Fe3nm%2F9WmsNzYwh%2FjlrGDMBUGA1UdIAQOMAwwCgYIKwYBBAHqAGQwagYDVR0fBGMwYTBfoF2gW4ZZaHR0cHM6Ly9jcmwtdWF0LWVjY2lkLmNjaWQubmV1c3Rhci5iaXovTmV1c3RhckVudGVycHJpc2VDZXJ0aWZpZWRDYWxsZXJJZEludGVybWVkaWF0ZS5jcmwwHQYDVR0OBBYEFIxKwMWpOLvYVK%2BavmzFmf94Eh6aMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA8ZgsiEJGHgvRrcD%2FQzzOnEjDkMSC0Sa2qwp5r4qOfq8CIF4bfiK%2F04iSrh2KKnf4A5NSRumFj58rLm3LjC8i%2BaJP)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN PNCB', but common name is 'STIR ECC PNCB' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.13568.100. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'STIR ECC PNCB' does not contain 'SHAKEN' |


Generated: 18 Aug 25 21:13 UTC