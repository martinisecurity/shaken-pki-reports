# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 762 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=AcmeTelecom\\, Inc., L=Somewhere, ST=VA, C=US\
Issuer: CN=SHAKEN, OU=VOIP, O=BigFatTelecom Corp., L=Someplace, ST=BC, C=CA\
Link: https://65.108.80.93/cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIB8zCCAZigAwIBAgIUU3UicVinFNZXYWN7WQQZuSbSOz4wCgYIKoZIzj0EAwIwbDELMAkGA1UEBhMCQ0ExCzAJBgNVBAgMAkJDMRIwEAYDVQQHDAlTb21lcGxhY2UxHDAaBgNVBAoME0JpZ0ZhdFRlbGVjb20gQ29ycC4xDTALBgNVBAsMBFZPSVAxDzANBgNVBAMMBlNIQUtFTjAeFw0yMzA5MTMxMzAxNDBaFw0yNTEyMTYxMzAxNDBaMGoxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJWQTESMBAGA1UEBwwJU29tZXdoZXJlMRowGAYDVQQKDBFBY21lVGVsZWNvbSwgSW5jLjENMAsGA1UECwwEVk9JUDEPMA0GA1UEAwwGU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3i6FwghqDsbq8fvpXsP%2BzzQyPMiCn4ZpB2EIYOpcvLu2e8wAsgzRr%2Bkro9TLl9%2FrBta8BO%2BaIP9Ywy95Fu8376MaMBgwFgYIKwYBBQUHARoECjAIoAYWBDEyMzQwCgYIKoZIzj0EAwIDSQAwRgIhAKdUP6ce0SlELQSJ%2Bzsnr8wa%2FaKsaDqzxyDEOIivFCpMAiEArkUAHaFsMsPskU4iIhxtg27eLTS7J7KSdHrhGJvD%2FIk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1234' |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_atis_basic_constraints](../../ISSUES/e_atis_basic_constraints/README.md) | error | ATIS1000080 | STI certificates shall contain a BasicConstraints extension marked critical |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | STI certificates shall contain a Key Usage extension marked as critical |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_authority_key_identifier](../../ISSUES/e_atis_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |


Generated: 15 Nov 23 18:10 UTC