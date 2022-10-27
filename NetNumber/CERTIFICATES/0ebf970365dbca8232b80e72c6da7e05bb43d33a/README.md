# STIR/SHAKEN CA Ecosystem Compliance
## NetNumber

### Certificate 0ebf970365dbca8232b80e72c6da7e05bb43d33a
Tested At: 2022-10-27 18:24:40 +0000 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 225 day(s)\
Subject: L=Austin, ST=Texas, O=Plivo Inc, C=US, CN=Plivo Inc\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1

Link: https://certificate.zt.plivo.com/cert09062023.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2zCCAmGgAwIBAgIJAKXEEo%2BeW3ZoMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjA2MTAwMDAwMDBaFw0yMzA2MDkwMDAwMDBaMFYxEjAQBgNVBAMMCVBsaXZvIEluYzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCVBsaXZvIEluYzEOMAwGA1UECAwFVGV4YXMxDzANBgNVBAcMBkF1c3RpbjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC%2FufDiYB7D220bJjcWjotKDZZ45eQrxiKG%2Bo0EMRip%2BHxYxbEc3XQ9AcU2vm6TdOcydRa0099nsFXGYWda8cnajgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDgwMEowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFNYHf0t1i%2BNxjuXpH1IEEuhcR42hMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwGIDwTWwe7d1VYWbFVmsUp55uDNBac3XmXq9nOrgwmYraUKsCtU6og0lKhNbAcfbgAjEA%2F8s8Lu7WqmW2qmM%2BRkycCf3cyKQtUklTG6P2v%2BJRXqBpTCEeUJ6MsUW3bSWIXm5T)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 800J' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| n_sti_certificate_policy_critical | notice | ATIS-1000080v4 | STI certificates should contain a CertificatePolicies extension marked uncritical |

Generated: 27/10/2022 at 18:24:52