# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 303d898eeeb3da9a375ea057a8f1fb1a00921e3f
Tested At: 2022-10-26 20:59:44 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/ff4f3455-b2b0-4f18-9435-ef5fa9a58b9f.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVjCCAfygAwIBAgIHXSkobwDwKzAKBggqhkjOPQQDAjBfMQswCQYDVQQGEwJVUzEVMBMGA1UECAwMUGVubnN5bHZhbmlhMRAwDgYDVQQKDAdDb21jYXN0MScwJQYDVQQDDB5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjIxMDA2MTAwNTQ5WhcNMjIxMTA1MTAwNTQ5WjBeMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxDzANBgNVBAMTBlNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCqK6ieSk7PZ2ga6AK%2BRaO68cfG0Ypff5hZmolmD93mAAUv%2FiP0zKN%2F2K2%2BKoPLGxOsBdKO17M1SWLGBP75kLR2jgaMwgaAwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUv44cYFUkxRFIz0JFRL%2Bw9Dor5igwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQ3NjEwMAoGCCqGSM49BAMCA0gAMEUCIQD8pfvMQkDVor9SvDVpjjRbWVNcQ80PuI4ujO9V9TMW%2BQIgbNlf%2BN8VpQAT0VPaW%2FrLEoSaOaG%2F74ZK1sqUpiDPuCI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_serial_number | error | ATIS-1000080v4 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

Generated: 26/10/2022 at 21:01:13