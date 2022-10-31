# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Inventive Labs Corp 649J

Tested At: 31 Oct 22 18:17 UTC\
Initial Validity Period: 180 day(s)\
Remaining Validity Period: 160 day(s)\
Subject: CN=SHAKEN Inventive Labs Corp 649J, OU=NOC, O=Inventive Labs Corp, ST=Colorado, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Inventive_Labs_Corp_649J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2jCCAoCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTyswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTk1NFoXDTIzMDQwOTE3MTk1NFowdjELMAkGA1UEBhMCVVMxETAPBgNVBAgMCENvbG9yYWRvMRwwGgYDVQQKDBNJbnZlbnRpdmUgTGFicyBDb3JwMQwwCgYDVQQLDANOT0MxKDAmBgNVBAMMH1NIQUtFTiBJbnZlbnRpdmUgTGFicyBDb3JwIDY0OUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXC%2Btlx4n9sblIwd5KcHwze6W96362PJik1R3HTDfaDPNhlzzu559LEsVbf0EXUI0lFybUJP9evOvhRIkKOec7o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NDlKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUW8sdf%2Bb2wURBqurmWvtsEfdWyQYwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIEf2TOjEKOuLimSZWLGwCQRf7dcyKtJIZtzPNsGO6utIAiEAj2PtdFXzVBIRmiv9yfylo3Rtn2%2Fqn08fZ6KdhG1ktGs%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 649J' |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31/10/2022 at 18:25:03