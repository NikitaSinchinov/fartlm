### README.md

---

# **Setup Instructions**

1. **Install Python 3.10**  
   Download and install from [python.org](https://www.python.org/downloads/).

2. **Create and Activate Virtual Environment**
   ```bash
   python3.10 -m venv venv
   source venv/bin/activate  # For Linux/macOS
   venv\Scripts\activate     # For Windows
   ```

3. **Install Dependencies**
   ```bash
   pip install rvc-python
   pip install torch==2.1.1+cu118 torchaudio==2.1.1+cu118 --index-url https://download.pytorch.org/whl/cu118
   ```

4. **Add Models**  
   Place the required model files in the `models/fart_lm` directory. The following files are required:
   - `.index` file (e.g., `model_name.index`)
   - `.pth` file (e.g., `model_weights.pth`)