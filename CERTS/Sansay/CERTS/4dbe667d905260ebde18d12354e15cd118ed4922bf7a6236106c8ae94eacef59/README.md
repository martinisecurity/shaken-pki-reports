# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Phone.com, Inc. 633J

Tested At: 31 Oct 22 19:21 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 26 day(s)\
Subject: CN=SHAKEN Phone.com\\, Inc. 633J, OU=voipteam, O=Phone.com\\, Inc., ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Phonedotcom_633J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2TCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3wwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzc0NFoXDTIyMTEyNjAwMzc0NFowdTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGDAWBgNVBAoMD1Bob25lLmNvbSwgSW5jLjERMA8GA1UECwwIdm9pcHRlYW0xJDAiBgNVBAMMG1NIQUtFTiBQaG9uZS5jb20sIEluYy4gNjMzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJTLc%2BXG1A5e1KyUnjizH7mo5f%2F%2B26dXea1HU0AeWM5RZHqKQFPEVumydDsK204FUayisWfuxa8XQh2AlfbaMV2jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDYzM0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTWuKTd71%2FET2TtuoPYJdJ4kF9f3TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAJa1KvLI1gMHgsJ8n8XSf7S9ooGQboKmRjPIJukDuVEjAiB3nC3arx8pt9lcCF2d2qYSFIElIz9FDxe1ptveVmJ7cA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 633J' |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 19:21:49