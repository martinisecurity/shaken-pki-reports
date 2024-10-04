# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Socket Telecom LLC 554a

Tested At: 04 Oct 24 16:05 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -7 day(s)\
Subject: CN=SHAKEN Socket Telecom LLC 554a, O=Socket Telecom LLC, ST=Missouri, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/554a/429C7C70711E3820F0B8E1DEAE6FF3262264A849.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAnCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkqEkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDgyNzE5NDkzNVoXDTI0MDkyNjE5NDkzNVowZjELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE1pc3NvdXJpMRswGQYDVQQKDBJTb2NrZXQgVGVsZWNvbSBMTEMxJzAlBgNVBAMMHlNIQUtFTiBTb2NrZXQgVGVsZWNvbSBMTEMgNTU0YTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJVzNFwGlHsxP3OQ7SqTs35lJw8Dkpv8MuvRgqt0kyrR4AXno9%2FeRzyx36RDh1fC8jK%2BiHSt4Vh2hIW4RdU1zVKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU1NGEwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMB0GA1UdDgQWBBRWElWJAtp3FKmg4TeSVyQW%2BjwfbjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIhAJfS%2BcHIfV9%2Fc%2B1965m74lu4RjGUdvTTWgn%2BtMpUQtZaAh9nUyPetywBcRfTH%2FHmmVmNjUgbB4E23LjpR6XDQGY4)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '554a' contains characters other than uppercase letters and numbers |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 554a', but common name is 'SHAKEN Socket Telecom LLC 554a' |


Generated: 04 Oct 24 16:29 UTC