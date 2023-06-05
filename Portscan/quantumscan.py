from qiskit import QuantumCircuit, execute, Aer
from scapy.all import IP, send

def scan_quantico(ip_destino, portas):
    num_qubits = 4

    circuito = QuantumCircuit(num_qubits, num_qubits)

    for i in range(num_qubits):
        circuito.h(i)

    for i in range(num_qubits - 1):
        circuito.cx(i, num_qubits - 1)

    circuito.measure(range(num_qubits), range(num_qubits))

    simulador = Aer.get_backend('qasm_simulator')
    job = execute(circuito, simulador, shots=1)
    resultado = job.result()

    contagens = resultado.get_counts(circuito)

    portas_abertas = []
    for estado, contagem in contagens.items():
        if estado[num_qubits - 1] == '1':
            porta_quantica = int(estado[:num_qubits - 1], 2)
            porta_real = portas[porta_quantica]
            portas_abertas.append(porta_real)

    return portas_abertas

def enviar_pacote_disfarcado(ip_destino):
    pacote_disfarcado = IP(dst=ip_destino) / "b'\x18U\x08\xac\x07\xf7\x01?\x1bw2P\x00)\x05r\xaaa\x99py\x7f\x9f\xac\xd8\xfe\x07\xfc1\x12Q-'"

    send(pacote_disfarcado)

    print("Anonymous packet successfully sent to", ip_destino)

def main():
    print("     .  _  .")
    print("    | | | | |")
    print("    | |-+-| |")
    print("  .-| | | |-.")
    print("  | | |_|_| |")
    print("  `-|       |-'")
    print("    |       |")
    print("    |       |")
    print("    `-------'")

    ip_destino = "10.100.4.4"
    portas = [22, 80, 443, 3389, 8080, 8000, 21, 3306]

    portas_abertas = scan_quantico(ip_destino, portas)

    if portas_abertas:
        for porta in portas_abertas:
            print("Port", porta, "open")
    else:
        print("No open ports found.")

    enviar_pacote_disfarcado(ip_destino)

if __name__ == "__main__":
    main()
