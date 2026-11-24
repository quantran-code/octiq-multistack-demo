import json
import socket
import threading
import urllib.error
import urllib.request
from collections.abc import Iterator
from http.server import ThreadingHTTPServer

import pytest

from demo.health import HEALTH_PATH, make_server


def _wait_until_accepting(host: str, port: int) -> None:
    """Block until the server thread is actually accepting connections."""
    for _ in range(100):
        try:
            with socket.create_connection((host, port), timeout=0.1):
                return
        except OSError:
            continue
    raise RuntimeError(f"server did not start listening on {host}:{port}")


@pytest.fixture
def health_server() -> Iterator[str]:
    server: ThreadingHTTPServer = make_server("127.0.0.1", 0)
    host, port = server.server_address[0], server.server_address[1]
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()
    _wait_until_accepting(host, port)
    try:
        yield f"http://{host}:{port}"
    finally:
        server.shutdown()
        server.server_close()
        thread.join(timeout=5)


def test_health_returns_ok(health_server: str) -> None:
    with urllib.request.urlopen(health_server + HEALTH_PATH) as response:
        assert response.status == 200
        assert response.headers["Content-Type"].startswith("application/json")
        body = response.read()

    assert json.loads(body) == {"status": "ok"}


def test_unknown_path_returns_404(health_server: str) -> None:
    with pytest.raises(urllib.error.HTTPError) as exc_info:
        urllib.request.urlopen(health_server + "/not-health")

    try:
        assert exc_info.value.code == 404
    finally:
        exc_info.value.close()
