# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 21e6797a2a553ddf3e1d5c9d5ceb99f93fa6c773
Tested At: 2022-10-27 22:31:47 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 157C, OU=SHAKEN, O=DigitalSpeed Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/82d97ecb-9b38-4926-9062-8e21b8986930/4a5ab5a3b56a2bb35ff2b29c212a2ab8.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FTCCAqOgAwIBAgIQUc%2FTIDaxT%2F2Jw1yE1UykHjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjQyMDUxNDdaFw0yMjEwMzEyMDUxNDZaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtEaWdpdGFsU3BlZWQgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE1N0MwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQtQtb5sgxAvrvZkWelm1wCwBrPmgw%2FAAKpLWW9rvdcQ1IaC%2B9s6jGF%2Fu399EDAixu%2BDrRE7GxN9N1E5mWvQMkco4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFCeG6sRn%2F%2F3KU%2BjOkGVue5T3Lys6MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE1N0MwCgYIKoZIzj0EAwIDSAAwRQIgMg5zgnvjZOm6AOGf9b%2FRhECe%2Bs1g5UAIQpa6p%2Br3T5oCIQCSGvVyToWMomZcZ1L6jsL1eChZ2MupKh1Sa2Qss13P2Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 22:33:03