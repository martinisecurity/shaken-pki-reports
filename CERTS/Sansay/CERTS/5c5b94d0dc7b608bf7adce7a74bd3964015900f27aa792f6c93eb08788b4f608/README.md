# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Vitelglobal communications 698K

Tested At: 22 Aug 24 15:32 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -190 day(s)\
Subject: CN=SHAKEN Vitelglobal communications 698K, OU=VitelGlobal, O=Vitelglobal communications, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Vitelglobal_communications_698K

[View certificate details](https://x509.io/?cert=MIIC8jCCApmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjWowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDExNTAwMDAwMFoXDTI0MDIxNDAwMDAwMFowgY4xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApOZXcgSmVyc2V5MSMwIQYDVQQKDBpWaXRlbGdsb2JhbCBjb21tdW5pY2F0aW9uczEUMBIGA1UECwwLVml0ZWxHbG9iYWwxLzAtBgNVBAMMJlNIQUtFTiBWaXRlbGdsb2JhbCBjb21tdW5pY2F0aW9ucyA2OThLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvtt4CnqggF7N89D4dz6ihm4kn7P0Icmdh5hHOmNK6HZrmGFDwBLxknC6qMoXC6RUc3F4wbp7h9TW7rFNQPJD%2B6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENjk4SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFIjqNILGMKF4i2EcHf7sfiCrx2DqMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBVJEb5PcKhFeysirZjdqNzAvWAitqHrUNXrv0dlkv9ZQIgLHjGpeK8bAEXSsleiYWFc%2B3%2F6MTpbH7HJmOKjQdYT5s%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 698K', but common name is 'SHAKEN Vitelglobal communications 698K' |


Generated: 22 Aug 24 15:44 UTC