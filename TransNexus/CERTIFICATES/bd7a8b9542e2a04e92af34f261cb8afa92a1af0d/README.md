# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate bd7a8b9542e2a04e92af34f261cb8afa92a1af0d
Tested At: 2022-10-27 21:42:42 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/4d384937fa4057d32f2f1aa0fa855ff4.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIQY6EwxAcQdPr59bd0iU3QFDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDI0MDVaFw0yMjEwMzAyMDI0MDRaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABN%2F9Qf3tRqpQ%2B6EGa%2BwFpFYKFtfpCbpSQzWzh9vYGfIXU67%2FuP9Zfc7xeyuIQnQgGbQJ%2Bsd5cuwHOdVG3S7auDGjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUafCCb%2B0bnwXQJWDMKM0URvm3k7YwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNIADBFAiEAmo3iwhkZlsczxJg04jPWy7XJQTQGVDt6pw1g%2B6f%2FFv4CIFimY7KQhQb3mdFbf%2BiQWRdktbXTbfpv7%2FfHBBEN2C12)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 21:42:52