# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Telxio Networks 492K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 155 day(s)\
Subject: CN=SHAKEN Telxio Networks 492K, OU=Telxio Networks, O=Telxio Networks, ST=Nevada, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/telxio-networks

[View certificate details](https://understandingwebpki.com/?cert=MIIC2zCCAoKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJka2wwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQxOTA2Mzk1N1oXDTI0MDQxODA2Mzk1N1oweDELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk5ldmFkYTEYMBYGA1UECgwPVGVseGlvIE5ldHdvcmtzMRgwFgYDVQQLDA9UZWx4aW8gTmV0d29ya3MxJDAiBgNVBAMMG1NIQUtFTiBUZWx4aW8gTmV0d29ya3MgNDkySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFzynS8V0YY1zCq4JkQYv7ex5ZXRYw71cML1gKMzP%2B2%2FA%2B2m2Kv5ztAJAmFHovWnYuO8Qp2NCrROnhpaYkJf7EWjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ5MkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTkqRhGcMI3L6DzihyzNma1WrVM7DAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgPblypSahYc9ZPUvctydyufbfP0c7q18BVJ2bBNpMvbYCICJCcC7F0BdR59OaiGbqR4Ocq6yQO2SCG4V3MbejgoMQ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 492K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 17:17 UTC