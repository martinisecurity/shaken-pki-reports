# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 27 Nov 23 22:21 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 750 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=AcmeTelecom\\, Inc., L=Somewhere, ST=VA, C=US\
Issuer: CN=SHAKEN, OU=VOIP, O=BigFatTelecom Corp., L=Someplace, ST=BC, C=CA\
Link: https://65.108.80.93/cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIB8zCCAZigAwIBAgIUU3UicVinFNZXYWN7WQQZuSbSOz4wCgYIKoZIzj0EAwIwbDELMAkGA1UEBhMCQ0ExCzAJBgNVBAgMAkJDMRIwEAYDVQQHDAlTb21lcGxhY2UxHDAaBgNVBAoME0JpZ0ZhdFRlbGVjb20gQ29ycC4xDTALBgNVBAsMBFZPSVAxDzANBgNVBAMMBlNIQUtFTjAeFw0yMzA5MTMxMzAxNDBaFw0yNTEyMTYxMzAxNDBaMGoxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJWQTESMBAGA1UEBwwJU29tZXdoZXJlMRowGAYDVQQKDBFBY21lVGVsZWNvbSwgSW5jLjENMAsGA1UECwwEVk9JUDEPMA0GA1UEAwwGU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3i6FwghqDsbq8fvpXsP%2BzzQyPMiCn4ZpB2EIYOpcvLu2e8wAsgzRr%2Bkro9TLl9%2FrBta8BO%2BaIP9Ywy95Fu8376MaMBgwFgYIKwYBBQUHARoECjAIoAYWBDEyMzQwCgYIKoZIzj0EAwIDSQAwRgIhAKdUP6ce0SlELQSJ%2Bzsnr8wa%2FaKsaDqzxyDEOIivFCpMAiEArkUAHaFsMsPskU4iIhxtg27eLTS7J7KSdHrhGJvD%2FIk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage](../../ISSUES/e_atis_ext_key_usage/README.md) | error | ATIS1000080 | Key Usage extension not found |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1234', but common name is 'SHAKEN' |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_ext_authority_key_identifier](../../ISSUES/e_atis_ext_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_atis_ext_basic_constraints](../../ISSUES/e_atis_ext_basic_constraints/README.md) | error | ATIS1000080 | BasicConstraints extension not found |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |


Generated: 27 Nov 23 22:56 UTC