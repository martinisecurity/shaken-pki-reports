# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 4a8b88b1dfa4c13c7cdfb12aa4ef2e3e9b21b340
Tested At: 2022-10-26 20:59:53 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN Mitel Cloud Services\\, Inc. 670J, OU=ComNet, O=Mitel Cloud Services\\, Inc., ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Mitel_Cloud_Services_Inc._670J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6jCCApGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT04wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE5MzAyMVoXDTIzMTAxMTE5MzAyMVowgYYxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMSMwIQYDVQQKDBpNaXRlbCBDbG91ZCBTZXJ2aWNlcywgSW5jLjEPMA0GA1UECwwGQ29tTmV0MS8wLQYDVQQDDCZTSEFLRU4gTWl0ZWwgQ2xvdWQgU2VydmljZXMsIEluYy4gNjcwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABClLDQj3bXHztq7%2Fkd%2BgsJS8fPZJvLaqKW3YrFJ2rwvtNMzKChOZ2AzfGvmaW%2Fhg9zD%2FvjYw7J2ZP4f%2FDXkUVyajgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDY3MEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQg4pDZbIkGDu%2Fkz8CqdvwPSWDlADAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgQXxNAwKEYI%2BYGHDh7GRPRMM6I8OwOQSquPG5ObOOhnICIDvXdhTOiWDKtT6X6l8XjfZXZhLqbG3wYDl9X7kGw6Wh)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 670J' |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 26/10/2022 at 21:01:13