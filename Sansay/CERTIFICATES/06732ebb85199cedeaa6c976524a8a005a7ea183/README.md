# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 06732ebb85199cedeaa6c976524a8a005a7ea183
Tested At: 2022-10-28 18:22:19 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: CN=SHAKEN Terra Nova Telecom 382G, OU=NOC, O=Terra Nova Telecom, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/382G_TERRANOVA_STIR_SHAKEN.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDhjCCAyugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjIyNTMzMloXDTIyMTEyNTIyNTMzMlowczELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGzAZBgNVBAoMElRlcnJhIE5vdmEgVGVsZWNvbTEMMAoGA1UECwwDTk9DMScwJQYDVQQDDB5TSEFLRU4gVGVycmEgTm92YSBUZWxlY29tIDM4MkcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQUZcBpTXzq%2F4l2CuZ90OBlLvIaqacu478cFveXc0%2Bt6sPgddYepZCyT9JcZG0NYLjUHUjCQ9WZslg8AbwbS0Qeo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDM4MkcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQZ3CGz9usyXL1owkd%2FQ8Oe1gPMCjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQC46hQH7%2BvhJAhIcGmPw2YnjJlj2eNQ36OukTpPCW%2Bo7QIhAIHlGuXewti84PDzV%2BLRJTqO5tMbE3HZydDQxz6zWdTR)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 382G' |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 28/10/2022 at 18:22:55