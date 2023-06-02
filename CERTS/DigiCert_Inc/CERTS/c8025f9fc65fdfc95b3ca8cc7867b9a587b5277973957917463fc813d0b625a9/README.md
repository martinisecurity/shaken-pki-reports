# STIR/SHAKEN CA Ecosystem Compliance

## Certificate DigiCert Global G2 TLS RSA SHA256 2020 CA1

Tested At: 02 Jun 23 01:12 UTC\
Initial Validity Period: 3652 day(s)\
Remaining Validity Period: 2858 day(s)\
Subject: CN=DigiCert Global G2 TLS RSA SHA256 2020 CA1, O=DigiCert Inc, C=US\
Issuer: CN=DigiCert Global Root G2, OU=www.digicert.com, O=DigiCert Inc, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIEyDCCA7CgAwIBAgIQDPW9BitWAvR6uFAsI8zwZjANBgkqhkiG9w0BAQsFADBhMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBHMjAeFw0yMTAzMzAwMDAwMDBaFw0zMTAzMjkyMzU5NTlaMFkxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxMzAxBgNVBAMTKkRpZ2lDZXJ0IEdsb2JhbCBHMiBUTFMgUlNBIFNIQTI1NiAyMDIwIENBMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMz3EGJPprtjb%2B2QUlbFbSd7ehJWivH0%2Bdbn4Y%2B9lavyYEEVcNsSAPonCrVXOFt9slGTcZUOakGUWzUb%2Bnv6u8W%2BJDD%2BVu%2FE832X4xT1FE3LpxDyFuqrIvAxIhFhaZAmunjZlx%2FjfWardUSVc8is%2F%2B9dCopZQ%2BGssjoP80j812s3wWPc3kbW20X%2BfSP9kOhRBx5Ro1%2FtSUZUfyyIxfQTnJcVPAPooTncaQwywa8WV0yUR0J8osicfebUTVSvQpmowQTCd5zWSOTOEeAqgJnwQ3DPP3Zr0UxJqyRewg2C%2FUaoq2yTzGJSQnWS%2BJr6Xl6ysGHlHx%2B5fwmY6D36g39HaaECAwEAAaOCAYIwggF%2BMBIGA1UdEwEB%2FwQIMAYBAf8CAQAwHQYDVR0OBBYEFHSFgMBmx9833s%2B9KTeqAx2%2B7c0XMB8GA1UdIwQYMBaAFE4iVCAYlebjbuYP%2Bvq5Eu0GF485MA4GA1UdDwEB%2FwQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwdgYIKwYBBQUHAQEEajBoMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wQAYIKwYBBQUHMAKGNGh0dHA6Ly9jYWNlcnRzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RHMi5jcnQwQgYDVR0fBDswOTA3oDWgM4YxaHR0cDovL2NybDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0R2xvYmFsUm9vdEcyLmNybDA9BgNVHSAENjA0MAsGCWCGSAGG%2FWwCATAHBgVngQwBATAIBgZngQwBAgEwCAYGZ4EMAQICMAgGBmeBDAECAzANBgkqhkiG9w0BAQsFAAOCAQEAkPFwyyiXaZd8dP3A%2BiZ7U6utzWX9upwGnIrXWkOH7U1MVl%2BtwcW1BSAuWdH%2FSvWgKtiwla3JLko716f2b4gp%2FDA%2FJIS7w7d7kwcsr4drdjPtAFVSslme5LnQ89%2FnD%2F7d%2BMS5EHKBCQRfz5eeLjJ1js%2BaWNJXMX43AYGyZm0pGrFmCW3RbpD0ufovARTFXFZkAdl9h6g4U5%2BLXUZtXMYnhIHUfoyMo5tS58aI7Dd8KvvwVVo4chDYABPPTHPbqjc1qCmBaZx2vN4Ye5DUys%2FvZwP9BFohFrH%2F6j%2Ff3IL16%2FRZkiMNJCqVJUzKoZHm1Lesh3Sz8W2jmdv51b2EQJ8HmA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ca_signature_algorithm](../../ISSUES/e_atis_ca_signature_algorithm/README.md) | error | ATIS1000080 | STI certificates shall contain a Signature Algorithm field with the value 'ecdsa-with-SHA256' |
| [e_atis_ca_subject_public_key](../../ISSUES/e_atis_ca_subject_public_key/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Public Key Info field specifying a Public Key Algorithm of "id-ecPublicKey" and containing a 256-bit public key |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Jun 23 01:12 UTC