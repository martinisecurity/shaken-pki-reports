# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

Name: e_atis_subject_cn_spc\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: For end-entity certificate, the Common Name attribute shall contain the text string SHAKEN, followed by a single space, followed by the SPC value identified in the TNAuthList of the end-entity certificate.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | Baltimore-Washington Telephone Company SHAKEN cert 8697 | [view](../../CERTS/6d933f79b12eb44c22c5fef74073e054fa5c31aa25a75811d8a032f3d894ab24/README.md) | Common name shall contain the text string 'SHAKEN 8697', but common name is 'Baltimore-Washington Telephone Company SHAKEN cert 8697' |
| error | Plivo Inc | [view](../../CERTS/fed50200daa631dd0cd7b74969c780f8d456dfd31db156c6cbea276f5a9a4cbf/README.md) | Common name shall contain the text string 'SHAKEN 800J', but common name is 'Plivo Inc' |
| error | DISH Wireless L.L.C.SHAKEN.490J | [view](../../CERTS/2943713c56f0705ed027ecffced5eb89cb1c36bb5386bdc36a6b8e5618ca2c9c/README.md) | Common name shall contain the text string 'SHAKEN 490J', but common name is 'DISH Wireless L.L.C.SHAKEN.490J' |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 05 Apr 24 19:04 UTC