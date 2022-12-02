# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Baltimore-Washington Telephone Company SHAKEN cert 8697

Tested At: 02 Dec 22 07:44 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 22 day(s)\
Subject: O=Baltimore-Washington Telephone Company, C=US, CN=Baltimore-Washington Telephone Company SHAKEN cert 8697\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://bwt.io/8697a2.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICvTCCAkKgAwIBAgIIKgjiTEpu8EgwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIyMTEyMzE3MDAwMFoXDTIyMTIyMzE2NTk1OVowgYAxQDA%2BBgNVBAMMN0JhbHRpbW9yZS1XYXNoaW5ndG9uIFRlbGVwaG9uZSBDb21wYW55IFNIQUtFTiBjZXJ0IDg2OTcxCzAJBgNVBAYTAlVTMS8wLQYDVQQKDCZCYWx0aW1vcmUtV2FzaGluZ3RvbiBUZWxlcGhvbmUgQ29tcGFueTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHs%2BqCy1MnQamE65Rvrgw8mRZZfo1383P2RjOmP6O0NM8g%2BTYQwDMYJI00EW5EXShu6KlCQl%2FqZhgdFk9KRL6oGjgZUwgZIwFgYIKwYBBQUHARoECjAIoAYWBDg2OTcwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFLwIUZBTUpt0TvFxtUhQdGvg09RFMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNpADBmAjEAs0%2B5GQBdgERP9zycrB1G0wC2z3w52n2%2FeaupUWeNmkDQDyxNSvalDwNLL9N3xy9pAjEAiEJIjENfRBRRRApULYuPFYM3zOR3XU15GI0OXu82bFGIB4%2FZ72zdixNVBH5q4E9i)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8697' |


Generated: 02 Dec 22 07:46 UTC