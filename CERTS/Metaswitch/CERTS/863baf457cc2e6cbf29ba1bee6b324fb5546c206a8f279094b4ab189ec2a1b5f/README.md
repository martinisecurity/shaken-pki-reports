# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Twin Lakes SHAKEN Cert 0579

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 866 day(s)\
Subject: CN=Twin Lakes SHAKEN Cert 0579, O=Twin Lakes, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/230562ee4f78391a4f429776b840b730947bb2f6

[View certificate details](https://understandingwebpki.com/?cert=MIICUDCCAfWgAwIBAgIQGP4VldCe2mCaLRZ0thlAHTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMzMTE2MDgxMloXDTI2MDMzMDE2MDgxMlowSDELMAkGA1UEBhMCVVMxEzARBgNVBAoMClR3aW4gTGFrZXMxJDAiBgNVBAMMG1R3aW4gTGFrZXMgU0hBS0VOIENlcnQgMDU3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC0k%2FH700uD%2FK8kkPi3jPMZYIk6T16Zg0nqzTnbjGXBv2krQiidZEffwf95s8rX2Qvgy1SVXHC2c%2FvjPBISXOBKjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA1NzkwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUw4A9y9Rf11VH3Fz4e3z0gr0FSGswHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhALAPEYzP9%2FRCKhNRJ118FYHPB3tSmoIbi0jBUnnj55NyAiEAnwc6kpjsxg05l0zcySP3hBSDP18mhifCudjnz4btj%2F8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0579' |


Generated: 15 Nov 23 17:17 UTC