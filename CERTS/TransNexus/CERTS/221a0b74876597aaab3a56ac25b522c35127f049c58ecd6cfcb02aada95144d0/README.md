# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TCN SHAKEN 706J

Tested At: 28 Apr 23 02:05 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -272 day(s)\
Subject: CN=TCN SHAKEN 706J, OU=TCN, O=TransNexus, C=US\
Issuer: CN=TransNexus Issuing CA G2, OU=Certification Authorities, O=TransNexus, C=US\
Link: https://certificates.transnexus.com/706J/9a9ea71d-1c8c-424c-8109-df524a4633d8.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICjTCCAjOgAwIBAgIQQJPGvayxdZg8r0JDxLe5qTAKBggqhkjOPQQDAjBpMQswCQYDVQQGEwJVUzETMBEGA1UEChMKVHJhbnNOZXh1czEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEhMB8GA1UEAxMYVHJhbnNOZXh1cyBJc3N1aW5nIENBIEcyMB4XDTIxMDcyOTIyMDQwNVoXDTIyMDcyOTIyMDQwNFowSjELMAkGA1UEBhMCVVMxEzARBgNVBAoMClRyYW5zTmV4dXMxDDAKBgNVBAsMA1RDTjEYMBYGA1UEAwwPVENOIFNIQUtFTiA3MDZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEW9%2BypyADHeBGNkLhqbdUTZSDEzPJYqkxotU6%2FcOM2UtI3qbwpC2%2FVoMuiD76aXLYG%2F7RyhYIjkZqm2b53ElZT6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQU6Thyv0Cp7TCpTkcmG%2ByM4ai%2FGHIwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAfBgNVHSMEGDAWgBT%2BJW42e1PDKhdzCELBZMZ5i6cmqDAWBggrBgEFBQcBGgQKMAigBhYENzA2SjAKBggqhkjOPQQDAgNIADBFAiEAtP%2FwcOE1zb5NtcmG7NzsTGuvIZWUVEqIFaHo7OzCoNMCIBHgJond3PlxN7zH9mp5IaVcOWpotFA4JfgOIzAE17%2FS)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Apr 23 02:17 UTC