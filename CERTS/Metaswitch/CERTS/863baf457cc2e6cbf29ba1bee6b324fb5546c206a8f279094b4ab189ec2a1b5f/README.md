# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Twin Lakes SHAKEN Cert 0579

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 777 day(s)\
Subject: CN=Twin Lakes SHAKEN Cert 0579, O=Twin Lakes, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/230562ee4f78391a4f429776b840b730947bb2f6

[View certificate details](https://understandingwebpki.com/?cert=MIICUDCCAfWgAwIBAgIQGP4VldCe2mCaLRZ0thlAHTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMzMTE2MDgxMloXDTI2MDMzMDE2MDgxMlowSDELMAkGA1UEBhMCVVMxEzARBgNVBAoMClR3aW4gTGFrZXMxJDAiBgNVBAMMG1R3aW4gTGFrZXMgU0hBS0VOIENlcnQgMDU3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC0k%2FH700uD%2FK8kkPi3jPMZYIk6T16Zg0nqzTnbjGXBv2krQiidZEffwf95s8rX2Qvgy1SVXHC2c%2FvjPBISXOBKjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA1NzkwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUw4A9y9Rf11VH3Fz4e3z0gr0FSGswHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhALAPEYzP9%2FRCKhNRJ118FYHPB3tSmoIbi0jBUnnj55NyAiEAnwc6kpjsxg05l0zcySP3hBSDP18mhifCudjnz4btj%2F8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0579', but common name is 'Twin Lakes SHAKEN Cert 0579' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 12 Feb 24 17:02 UTC