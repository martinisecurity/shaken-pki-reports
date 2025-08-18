# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Baltimore-Washington Telephone Company SHAKEN cert 8697

Tested At: 18 Aug 25 20:04 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 13 day(s)\
Subject: O=Baltimore-Washington Telephone Company, C=US, CN=Baltimore-Washington Telephone Company SHAKEN cert 8697\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://bwt.io/8697a2.crt

[View certificate details](https://x509.io/?cert=MIICvTCCAkKgAwIBAgIIdedWjRlNGYYwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTI1MDgwMTA3MDAwMFoXDTI1MDgzMTA3MDAwMFowgYAxQDA%2BBgNVBAMMN0JhbHRpbW9yZS1XYXNoaW5ndG9uIFRlbGVwaG9uZSBDb21wYW55IFNIQUtFTiBjZXJ0IDg2OTcxCzAJBgNVBAYTAlVTMS8wLQYDVQQKDCZCYWx0aW1vcmUtV2FzaGluZ3RvbiBUZWxlcGhvbmUgQ29tcGFueTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHs%2BqCy1MnQamE65Rvrgw8mRZZfo1383P2RjOmP6O0NM8g%2BTYQwDMYJI00EW5EXShu6KlCQl%2FqZhgdFk9KRL6oGjgZUwgZIwFgYIKwYBBQUHARoECjAIoAYWBDg2OTcwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFLwIUZBTUpt0TvFxtUhQdGvg09RFMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNpADBmAjEA8HaG%2FOVAQvzf%2FV%2FWYbdgO7q7E2Q0uXRmYnMlTK3hC9xBass%2FDJuzCvzhPRFqjKrbAjEAv996psKWWqsdhdUdftulaLuNDALj5tjOAxISxKJmxr%2FCOztLHTnuaUdd6fO4YGJs)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8697', but common name is 'Baltimore-Washington Telephone Company SHAKEN cert 8697' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |


Generated: 18 Aug 25 21:13 UTC