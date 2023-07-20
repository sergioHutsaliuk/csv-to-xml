package convertor

import "fmt"

func withVAT(item *DataCSV) string {
	template :=
		`<Factura>
			<Antet>
				<FurnizorNume>BV MEDIA CORPORATION S.R.L.</FurnizorNume>
				<FurnizorCIF>42485343</FurnizorCIF>
				<FurnizorNrRegCom>J35/1021/2020</FurnizorNrRegCom>
				<FurnizorCapital>200.00</FurnizorCapital>
				<FurnizorAdresa>Municipiul Timişoara, Strada 1 DECEMBRIE, Nr. 94, Ap. 1, Judet Timiş</FurnizorAdresa>
				<FurnizorBanca>LIBRA BANK S.A.</FurnizorBanca>
				<FurnizorIBAN>RO62BREL0005600852510100</FurnizorIBAN>
				<FurnizorInformatiiSuplimentare>Tel. 0212088088 LIBRA BANK S.A. RO62BREL0005600852510100</FurnizorInformatiiSuplimentare>
				<ClientNume>%s</ClientNume>
				<ClientInformatiiSuplimentare/>
				<ClientCIF/>
				<ClientNrRegCom/>
				<ClientJudet/>
				<ClientTara>%s</ClientTara>
				<ClientAdresa>%s</ClientAdresa>
				<ClientTelefon/>
				<ClientEmail>%s</ClientEmail>
				<ClientBanca/>
				<ClientIBAN/>
				<FacturaNumar>%s</FacturaNumar>
				<FacturaData>%s</FacturaData>
				<FacturaScadenta>%s</FacturaScadenta>
				<FacturaTaxareInversa>Nu</FacturaTaxareInversa>
				<FacturaTVAIncasare>Nu</FacturaTVAIncasare>
				<FacturaInformatiiSuplimentare/>
				<FacturaMoneda>%s</FacturaMoneda>
				<FacturaCotaTVA>%s</FacturaCotaTVA>
				<FacturaGreutate>0.000</FacturaGreutate>
				<FacturaAccize>0.00</FacturaAccize>
			</Antet>
			<Detalii>
				<Continut>
					<Linie>
						<LinieNrCrt>1</LinieNrCrt>
						<Descriere>VANZARI EU</Descriere>
						<CodArticolFurnizor/>
						<CodArticolClient/>
						<CodBare/>
						<InformatiiSuplimentare/>
						<UM>1</UM>
						<Cantitate>1.000</Cantitate>
						<Pret>%s</Pret>
						<Valoare>%s</Valoare>
						<ProcTVA>%s</ProcTVA>
						<TVA>%s</TVA>
					</Linie>
				</Continut>
				<txtObservatii1/>
			</Detalii>
			<Sumar>
				<TotalValoare>%s</TotalValoare>
				<TotalTVA>%s</TotalTVA>
				<Total>%s</Total>
				<LinkPlata/>
			</Sumar>
			<Observatii>
				<txtObservatii/>
				<SoldClient/>
				<ModalitatePlata/>
			</Observatii>
		</Factura>`

	return fmt.Sprintf(template, item.ClientName, item.ClientCountry, item.ClientAddress, item.ClientEmail, item.Number, item.Date, item.PaymentDate, item.Currency, item.VATFull, item.NetAmount+"00", item.NetAmount, item.VATShort, item.VATAmount, item.NetAmount, item.VATAmount, item.TotalAmount)
}

func withoutVAT(item *DataCSV) string {
	template :=
		`<Factura>
			<Antet>
				<FurnizorNume>BV MEDIA CORPORATION S.R.L.</FurnizorNume>
				<FurnizorCIF>42485343</FurnizorCIF>
				<FurnizorNrRegCom>J35/1021/2020</FurnizorNrRegCom>
				<FurnizorCapital>200.00</FurnizorCapital>
				<FurnizorAdresa>Municipiul Timişoara, Strada 1 DECEMBRIE, Nr. 94, Ap. 1, Judet Timiş</FurnizorAdresa>
				<FurnizorBanca>LIBRA BANK S.A.</FurnizorBanca>
				<FurnizorIBAN>RO62BREL0005600852510100</FurnizorIBAN>
				<FurnizorInformatiiSuplimentare>Tel. 0212088088 LIBRA BANK S.A. RO62BREL0005600852510100</FurnizorInformatiiSuplimentare>
				<ClientNume>%s</ClientNume>
				<ClientInformatiiSuplimentare/>
				<ClientCIF/>
				<ClientNrRegCom/>
				<ClientJudet/>
				<ClientTara>%s</ClientTara>
				<ClientAdresa>%s</ClientAdresa>
				<ClientTelefon/>
				<ClientEmail>%s</ClientEmail>
				<ClientBanca/>
				<ClientIBAN/>
				<FacturaNumar>%s</FacturaNumar>
				<FacturaData>%s</FacturaData>
				<FacturaScadenta>%s</FacturaScadenta>
				<FacturaTaxareInversa>Nu</FacturaTaxareInversa>
				<FacturaTVAIncasare>Nu</FacturaTVAIncasare>
				<FacturaInformatiiSuplimentare/>
				<FacturaMoneda>%s</FacturaMoneda>
				<FacturaGreutate>0.000</FacturaGreutate>
				<FacturaAccize>0.00</FacturaAccize>
			</Antet>
			<Detalii>
				<Continut>
					<Linie>
						<LinieNrCrt>1</LinieNrCrt>
						<Descriere>VANZARI NON EU</Descriere>
						<CodArticolFurnizor/>
						<CodArticolClient/>
						<CodBare/>
						<InformatiiSuplimentare/>
						<UM>1</UM>
						<Cantitate>1.000</Cantitate>
						<Pret>%s</Pret>
						<Valoare>%s</Valoare>
						<ProcTVA>0</ProcTVA>
						<TVA>0.00</TVA>
					</Linie>
				</Continut>
				<txtObservatii1/>
			</Detalii>
			<Sumar>
				<TotalValoare>%s</TotalValoare>
				<TotalTVA>0.00</TotalTVA>
				<Total>%s</Total>
				<LinkPlata/>
			</Sumar>
			<Observatii>
				<txtObservatii/>
				<SoldClient/>
				<ModalitatePlata/>
			</Observatii>
		</Factura>`

	return fmt.Sprintf(template, item.ClientName, item.ClientCountry, item.ClientAddress, item.ClientEmail, item.Number, item.Date, item.PaymentDate, item.Currency, item.NetAmount+"00", item.NetAmount, item.NetAmount, item.TotalAmount)
}
