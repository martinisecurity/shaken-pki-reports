# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Fonative, Inc. 684J

Tested At: 02 Dec 22 07:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 14 day(s)\
Subject: CN=SHAKEN Fonative\\, Inc. 684J, OU=Operations, O=Fonative\\, Inc., ST=Massachusetts, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Fonative_684J

[View certificate details](https://understandingwebpki.com/?cert=MIIC3TCCAoKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVYIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMDI1OVoXDTIyMTIxNTIxMDI1OVoweDELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDU1hc3NhY2h1c2V0dHMxFzAVBgNVBAoMDkZvbmF0aXZlLCBJbmMuMRMwEQYDVQQLDApPcGVyYXRpb25zMSMwIQYDVQQDDBpTSEFLRU4gRm9uYXRpdmUsIEluYy4gNjg0SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPrea997HVIgkzfIebrVo9iagsCwTM6hf23MV%2FQjwF4X%2BpoCGAJZvQ2j7pW%2B24cDtkRwWBeMx52rB5pAA87RP4ejgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDY4NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBS3E%2FxAvMmocKv0aoiX6LVbK2Si0jAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJEmMYO3NT%2FSxkN%2BLLfT2RlTeDKva99VptHI1i9o5QvdAiEApFUB3CKrKX3hQSMRForOn0hCCtFd1P9GSvhpfDs6S6M%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 684J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 02 Dec 22 07:46 UTC