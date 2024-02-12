# STIR/SHAKEN CA Ecosystem Compliance

## Certificate cert.stir.t-mobile.com

Tested At: 12 Feb 24 16:58 UTC\
Initial Validity Period: 366 day(s)\
Remaining Validity Period: -212 day(s)\
Subject: CN=cert.stir.t-mobile.com, OU=T-Mobile USA\\, Inc, O=T-Mobile USA\\, Inc., L=Bothell, ST=Washington, C=US\
Issuer: CN=TMOBILE-PROD-SUB-STIRSHAKEN-EC, O=TMOBILE-USA, C=US\
Link: https://t-mobile-sticr.fosrvt.com/88a8e33055e725475530660e5d6c40d6adbe37ab7ae0ecc64b50205629548ae9.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICyTCCAm%2BgAwIBAgIPAYH%2Bo0TdISlq5niitURDMAoGCCqGSM49BAMCMEwxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtUTU9CSUxFLVVTQTEnMCUGA1UEAxMeVE1PQklMRS1QUk9ELVNVQi1TVElSU0hBS0VOLUVDMB4XDTIyMDcxNDIxMDUyMVoXDTIzMDcxNDIxMzUyMVowgY4xCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdCb3RoZWxsMRswGQYDVQQKExJULU1vYmlsZSBVU0EsIEluYy4xGjAYBgNVBAsTEVQtTW9iaWxlIFVTQSwgSW5jMR8wHQYDVQQDExZjZXJ0LnN0aXIudC1tb2JpbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAER5y7ZAcpC5BgT4ElU782y9VvyzL5CSXR5dDg8QiLkbaMxuUJoFC6Yzdma2m4Jiz26UZhBA2AX1BDeGO29WDQZaOB8DCB7TAfBgNVHSMEGDAWgBQC2X9YhlA%2FeaB9iJtcgWc9pXwKeDAdBgNVHQ4EFgQUSnET9T3YyVih6PjLmbaCikNgEHswDgYDVR0PAQH%2FBAQDAgeAMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAhBgNVHREEGjAYghZjZXJ0LnN0aXIudC1tb2JpbGUuY29tMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENjUyOTAKBggqhkjOPQQDAgNIADBFAiBf03jL8qMfp5flDibYOien7k7E%2F9XlsSR0ReVxahdkWgIhAKvf9Zd8t%2BlIl6LdJ90f4MqUHtsHRQ%2FJVQ4Kty0IFM1h)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'cert.stir.t-mobile.com' does not contain 'SHAKEN' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 6529', but common name is 'cert.stir.t-mobile.com' |
| [e_atis_ext_basic_constraints](../../ISSUES/e_atis_ext_basic_constraints/README.md) | error | ATIS1000080 | BasicConstraints extension not found |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC