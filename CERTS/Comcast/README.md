# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 54 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 35 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 19 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 19 certificates expire in the next 30 days
- 19.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 19 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 19 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 50.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 6640 days is the average remaining validity for the certificates in the corpus
- 7300 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_crl_distribution_struct_ca](ISSUES/e_atis_ext_crl_distribution_struct_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_serial_number_size_ca](ISSUES/e_atis_serial_number_size_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 20&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/2102f404f08ead6467bfe6ba6a9482b20ad32e7ed0dd3d1bcde22fa9f76609d6/README.md) |
| 22&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 21&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/b31fbddcbe8e8e6e9defa4ecb019d7499987cbe51dc8f45ffd037ba6e797db5a/README.md) |
| 23&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 22&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/1de21698d9c145c20b1dd7437e553659b4dfc8134bf4dfc9fbdd2703f16b8837/README.md) |
| 24&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 23&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/ca1dc50d6b955112c1892558f27d303dabd429d903839bb2956e80c71689c8fd/README.md) |
| 25&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 24&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/9c3cb608e998096acfb0c08470e299d1648d7ff41875401270eb886785ff6cde/README.md) |
| 26&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 25&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/4f6d8b28a60487a06d12f19ccb993e2d27f705420f5d221656ecd3fbb3355fc4/README.md) |
| 28&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 27&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/9a86ee41a77e352f16f3f3fe841c4a029b7f5d1c67f168f1f6ac7d35b62f436e/README.md) |
| 29&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 28&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/6eb7de61b4e9b6c12cb841e72a2e08a76044c7842388306f4caf58e10597a57e/README.md) |
| 30&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 29&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/543f42442b0e50da8d0672ce11933ffeab7e107fbd66b59cf2b709cbf94c799d/README.md) |
| 31&#160;Jul&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 30&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/8436b9f915d6a782c9092ede4933cad91924351fed7b5cc947edde5414d5bed8/README.md) |
| 01&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 31&#160;Aug&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/87c9b6bc37659c556c667be856f1eb88cda0fd523a4ccef775dc22a7a475d965/README.md) |
| 02&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 01&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/9c42932d4cb49ee2b5a7ba1c2c3754e2ef7cf95f578ece82facbfef6e99d4dcb/README.md) |
| 04&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 03&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/7326a083b19feba962e4542b829208e4ea80211c079ce2b98e146a8b3f13a1f2/README.md) |
| 05&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 04&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/85b9eb43ef4dd8a8c7eb749910f53211552fa08cbe79e1b6b93cf914b542dde0/README.md) |
| 06&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 05&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/2b4a209439e359fea4b2b5f67843431f83baa4c64508ca711ad8d87fa1de4961/README.md) |
| 07&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 06&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/b9554281e5345e66c7492a1eaae9586371a3fd7ad7df2cb01bf74f1229303f03/README.md) |
| 08&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 07&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/3539e01043a28008b6645fd76212b0c26352f35bb64313315af40eb206979c02/README.md) |
| 09&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 08&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/26f813ca5d5ace912f6d1368071f068b3c9b256938bed85c308c5fddbc2a73bb/README.md) |
| 11&#160;Aug&#160;25&#160;07:12&#160;UTC | SHAKEN 318J | 10&#160;Sep&#160;25&#160;07:12&#160;UTC | true | [view](CERTS/4450507e05c0e0ce81c8bf32ee1410f8a3d51ca13f630b2b29b2cf201690f0ae/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 11&#160;Jan&#160;24&#160;02:59&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jan&#160;44&#160;02:59&#160;UTC | true | [view](CERTS/25af737667ed8b05cb8b8e7f44b2d7b5861551bea95ec48a73306ec75a92a662/README.md) |


Generated: 18 Aug 25 21:13 UTC