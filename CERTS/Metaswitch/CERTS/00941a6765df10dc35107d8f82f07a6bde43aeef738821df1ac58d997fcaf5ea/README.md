# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Cellular South Licenses SHAKEN Cert 6581

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 903 day(s)\
Subject: CN=Cellular South Licenses SHAKEN Cert 6581, O=Cellular South Licenses  DBA C Spire, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/d1af372b9cdb3ceba1f85f4ec3657797b3eb1c7f

[View certificate details](https://x509.io/?cert=MIICdjCCAhygAwIBAgIQCf4UUy21rSbF6CZBorxwszAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDIxNjEwMzI1N1oXDTI3MDIxNTEwMzI1N1owbzELMAkGA1UEBhMCVVMxLTArBgNVBAoMJENlbGx1bGFyIFNvdXRoIExpY2Vuc2VzICBEQkEgQyBTcGlyZTExMC8GA1UEAwwoQ2VsbHVsYXIgU291dGggTGljZW5zZXMgU0hBS0VOIENlcnQgNjU4MTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCAeLmZqXfIGFkvJwf4qeK9vxS7siRmkPwri7YDPb53yeCNgt3mWQwykPSL8RE0AaspuQheBWr09FnLHPSJyOJejgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDY1ODEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUH7COkc0FMyxmh8dcI%2BC25%2BHKEzEwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAMVfm5LoeZUkss7WmqeqrXLnlhbr3hSLGQXkFeIRfF1AAiB1C2KKg1Iah%2FpU8QWFbidHzR18uu6rFuQNpXRTGW8yUQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 6581', but common name is 'Cellular South Licenses SHAKEN Cert 6581' |


Generated: 26 Aug 24 18:03 UTC