# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Plivo Inc

Tested At: 21 Nov 22 20:43 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 200 day(s)\
Subject: L=Austin, ST=Texas, O=Plivo Inc, C=US, CN=Plivo Inc\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://certificate.zt.plivo.com/cert09062023.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC2zCCAmGgAwIBAgIJAKXEEo%2BeW3ZoMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjA2MTAwMDAwMDBaFw0yMzA2MDkwMDAwMDBaMFYxEjAQBgNVBAMMCVBsaXZvIEluYzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCVBsaXZvIEluYzEOMAwGA1UECAwFVGV4YXMxDzANBgNVBAcMBkF1c3RpbjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC%2FufDiYB7D220bJjcWjotKDZZ45eQrxiKG%2Bo0EMRip%2BHxYxbEc3XQ9AcU2vm6TdOcydRa0099nsFXGYWda8cnajgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDgwMEowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFNYHf0t1i%2BNxjuXpH1IEEuhcR42hMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwGIDwTWwe7d1VYWbFVmsUp55uDNBac3XmXq9nOrgwmYraUKsCtU6og0lKhNbAcfbgAjEA%2F8s8Lu7WqmW2qmM%2BRkycCf3cyKQtUklTG6P2v%2BJRXqBpTCEeUJ6MsUW3bSWIXm5T)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 800J' |


Generated: 21 Nov 22 20:55 UTC