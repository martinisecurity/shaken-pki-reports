# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Plivo Inc

Tested At: 28 Nov 23 19:52 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: -172 day(s)\
Subject: L=Austin, ST=Texas, O=Plivo Inc, C=US, CN=Plivo Inc\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://certificate.zt.plivo.com/cert09062023.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC2zCCAmGgAwIBAgIJAKXEEo%2BeW3ZoMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjA2MTAwMDAwMDBaFw0yMzA2MDkwMDAwMDBaMFYxEjAQBgNVBAMMCVBsaXZvIEluYzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCVBsaXZvIEluYzEOMAwGA1UECAwFVGV4YXMxDzANBgNVBAcMBkF1c3RpbjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC%2FufDiYB7D220bJjcWjotKDZZ45eQrxiKG%2Bo0EMRip%2BHxYxbEc3XQ9AcU2vm6TdOcydRa0099nsFXGYWda8cnajgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDgwMEowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFNYHf0t1i%2BNxjuXpH1IEEuhcR42hMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwGIDwTWwe7d1VYWbFVmsUp55uDNBac3XmXq9nOrgwmYraUKsCtU6og0lKhNbAcfbgAjEA%2F8s8Lu7WqmW2qmM%2BRkycCf3cyKQtUklTG6P2v%2BJRXqBpTCEeUJ6MsUW3bSWIXm5T)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Plivo Inc' does not contain 'SHAKEN' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 800J', but common name is 'Plivo Inc' |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC