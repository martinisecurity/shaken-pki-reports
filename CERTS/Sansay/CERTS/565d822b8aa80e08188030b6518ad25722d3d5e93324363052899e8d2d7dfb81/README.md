# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Touchtone 683A

Tested At: 26 Aug 24 17:48 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -51 day(s)\
Subject: CN=SHAKEN Touchtone 683A, OU=NOC, O=Touchtone, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/683A/429C7C70711E3820F0B8E1DEAE6FF32622649EEF.pem

[View certificate details](https://x509.io/?cert=MIICyDCCAm6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknu8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDYwNTIyNDcxN1oXDTI0MDcwNTIyNDcxN1owZDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxEjAQBgNVBAoMCVRvdWNodG9uZTEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gVG91Y2h0b25lIDY4M0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATwPwhwR259HmHeCt9A84ohVtuEKJYTrlFMI5HXCls08mFEPsyKfoW%2BySkQmaGCaCmCEHCD6rRfhou6X8dWmxB4o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODNBMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU4fP7CWiieHL3QK4YpzPbWW6r8hkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIH%2BRMQFNpAOEEhoBt%2BbZ9IJU9sWf%2BADxSznYan7jKsprAiEA75pu7TeYGIhji8Z5I9qQd9mq4wfcvVM3O37oV73Shok%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 683A', but common name is 'SHAKEN Touchtone 683A' |


Generated: 26 Aug 24 18:03 UTC