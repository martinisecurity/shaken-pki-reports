# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Baltimore-Washington Telephone Company SHAKEN cert 8697

Tested At: 12 Apr 23 01:01 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 12 day(s)\
Subject: O=Baltimore-Washington Telephone Company, C=US, CN=Baltimore-Washington Telephone Company SHAKEN cert 8697\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://bwt.io/8697a2.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICvDCCAkKgAwIBAgIIcgItUCtpLe0wCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMDMyMzE1MDAwMFoXDTIzMDQyMzE0NTk1OVowgYAxQDA%2BBgNVBAMMN0JhbHRpbW9yZS1XYXNoaW5ndG9uIFRlbGVwaG9uZSBDb21wYW55IFNIQUtFTiBjZXJ0IDg2OTcxCzAJBgNVBAYTAlVTMS8wLQYDVQQKDCZCYWx0aW1vcmUtV2FzaGluZ3RvbiBUZWxlcGhvbmUgQ29tcGFueTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHs%2BqCy1MnQamE65Rvrgw8mRZZfo1383P2RjOmP6O0NM8g%2BTYQwDMYJI00EW5EXShu6KlCQl%2FqZhgdFk9KRL6oGjgZUwgZIwFgYIKwYBBQUHARoECjAIoAYWBDg2OTcwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFLwIUZBTUpt0TvFxtUhQdGvg09RFMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNoADBlAjAz0XAmMbIavTZ9L8CSc7g1EpYsN4wsNmgMIdFnB4Bq1RDrswv9R%2FSG8pDDBkPEIQoCMQC%2Fe1j5u95VLKlvh2afmGKAjYTlt%2Fve%2BTiOVB5zj%2FvMFy8z2c%2Brm1JLMtiKWC%2F4CEg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8697' |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |


Generated: 12 Apr 23 01:46 UTC