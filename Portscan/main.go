package main

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func getServiceName(port int) string {
	_, err := net.LookupPort("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		return "unknown"
	}
	// Aqui você pode adicionar mapeamentos personalizados para serviços conhecidos,
	// se quiser exibir um nome específico em vez de apenas o número da porta.
	return fmt.Sprintf("port %d", port)
}

func scanPort(ipDestino string, porta int) (bool, string) {
	endereco := fmt.Sprintf("%s:%d", ipDestino, porta)
	conn, err := net.Dial("tcp", endereco)
	if err != nil {
		return false, ""
	}
	defer conn.Close()

	service := getServiceName(porta)

	return true, service
}

func enviarPacoteDisfarcado(ipDestino string) {
	// Construindo o pacote disfarçado
	pacote := gopacket.NewSerializeBuffer()
	opcoes := gopacket.SerializeOptions{}

	ip := layers.IPv4{
		SrcIP:    net.ParseIP("192.168.0.1"), // IP de origem
		DstIP:    net.ParseIP(ipDestino),     // IP de destino
		Protocol: layers.IPProtocolTCP,
	}

	tcp := layers.TCP{
		SrcPort: layers.TCPPort(1234), // Porta de origem
		DstPort: layers.TCPPort(80),   // Porta de destino
	}

	// Modificar os campos TCP para evasão de assinatura
	tcp.Checksum = 0           // Zerar o checksum para que seja recalculado automaticamente
	tcpWindowScale := uint8(0) // Valor aleatório para escala de janela TCP
	tcp.Options = []layers.TCPOption{
		{layers.TCPOptionKindWindowScale, 3, []byte{tcpWindowScale}},
		{layers.TCPOptionKindMSS, 4, []byte{0x05, 0xb4}}, // Tamanho máximo do segmento (MSS)
	}

	tcp.SetNetworkLayerForChecksum(&ip)

	// Adicionando os dados do pacote disfarçado
	dados := []byte("b'\x18U\x08\xac\x07\xf7\x01?\x1bw2P\x00)\x05r\xaaa\x99py\x7f\x9f\xac\xd8\xfe\x07\xfc1\x12Q-'")
	payload := gopacket.Payload(dados)

	// Adicionando as camadas ao pacote
	err := gopacket.SerializeLayers(pacote, opcoes,
		&ip,
		&tcp,
		payload,
	)

	if err != nil {
		fmt.Println("Erro ao construir o pacote disfarçado:", err)
		return
	}

	// Enviando o pacote disfarçado
	conn, err := net.Dial("ip4:tcp", ipDestino)
	if err != nil {
		fmt.Println("Erro ao enviar o pacote disfarçado:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(pacote.Bytes())
	if err != nil {
		fmt.Println("Erro ao enviar o pacote disfarçado:", err)
		return
	}

	fmt.Println("Pacote disfarçado enviado com sucesso para", ipDestino)
}

func main() {
	fmt.Println("     .  _  .")
	fmt.Println("    | | | | |")
	fmt.Println("    | |-+-| |")
	fmt.Println("  .-| | | |-.")
	fmt.Println("  | | |_|_| |")
	fmt.Println("  `-|       |-'")
	fmt.Println("    |       |")
	fmt.Println("    |       |")
	fmt.Println("    `-------'")

	ipDestino := "ofuscated"
	portas := []int{21, 22, 23, 25, 53, 67, 80, 107, 109, 110, 123, 123, 137, 138, 139, 161, 194, 220, 389, 411, 412, 445, 465, 513, 514, 3306, 3389, 366, 443, 465, 513, 901, 993, 995, 8888, 1024, 8000, 8443}

	portasAbertas := []int{} // Armazena as portas abertas encontradas

	for _, porta := range portas {
		if status, service := scanPort(ipDestino, porta); status {
			fmt.Println("Port", porta, "open - Service:", service)
			portasAbertas = append(portasAbertas, porta) // Adiciona a porta à lista de portas abertas
		} else {

		}
	}

	// Imprime apenas as portas abertas encontradas
	fmt.Println("Open ports:")
	for _, porta := range portasAbertas {
		fmt.Println(porta)
	}

	// Aguarda alguns segundos antes de enviar o pacote disfarçado
	time.Sleep(5 * time.Second)

	enviarPacoteDisfarcado(ipDestino)
}
