import logging
from logging.handlers import RotatingFileHandler
import os
import requests
from flask import Flask, request, render_template, stream_template, jsonify
from flask.logging import default_handler
from bs4 import BeautifulSoup

# Initialize Flask
app = Flask(__name__)

def configure_logging(app):
    # Logging Configuration
    if app.config['LOG_WITH_GUNICORN']:
        gunicorn_error_logger = logging.getLogger('gunicorn.error')
        app.logger.handlers.extend(gunicorn_error_logger.handlers)
        app.logger.setLevel(logging.DEBUG)
    else:
        file_handler = RotatingFileHandler('instance/flask-user-management.log',
                                           maxBytes=16384,
                                           backupCount=20)
        file_formatter = logging.Formatter('%(asctime)s %(levelname)s %(threadName)s-%(thread)d: %(message)s [in %(filename)s:%(lineno)d]')
        file_handler.setFormatter(file_formatter)
        file_handler.setLevel(logging.INFO)
        app.logger.addHandler(file_handler)

    # Remove the default logger configured by Flask
    app.logger.removeHandler(default_handler)

    app.logger.info('...Starting the Flask Server...')


# Set a secret key for session management
app.secret_key = os.urandom(32)

#Get the url to the API from environment variable
api_url = os.environ.get('API_URL')

# Function to check API status
def API_OK():
    try:
        response = requests.get(api_url)
        return response.status_code == 418
    except requests.ConnectionError:
        return False

# Specify Index route of the flask application
@app.route('/')
def index():
    # If API is down render website down html
    if not API_OK():
        return render_template('website_down.html')
    
    # Determine if user is on Mobile of Computer
    # Used to load specific CSS in the HTML templates
    user_agent = request.headers.get('User-Agent')
    if any(keyword in user_agent.lower() for keyword in ['iphone', 'android', 'blackberry', 'mobile', 'linux']):
        if 'ubuntu' in user_agent.lower():
            mobile = ""
        else:
            mobile = "true"
    else:
        mobile = ""

    # Get the API Query parameters
    if request.args.get('languages') != None:
        language = request.args.get('languages')
    else:
        language = ""

    if request.args.get('search') != None:
        search_query = request.args.get('search')
    else:
        search_query = ""

    if request.args.get('page') != None:
        page = request.args.get('page')
    else:
        page = ""

    # Fetch JSON data from the endpoints
    url = f'{api_url}/library/v1/?languages={language}&search={search_query}&page={page}'
    url_support_languages = f'{api_url}/librarystats/v1/supported_languages/'
    response = requests.get(url)
    response_support_languages = requests.get(url_support_languages)
    data = response.json()
    support_languages = response_support_languages.json()

    # Render HTML template with the fetched data
    return stream_template('index.html',
                            results=data['results'],
                            next=data['next'],
                            previous=data['previous'],
                            language=language,
                            search=search_query,
                            support_languages=support_languages,
                            mobile=mobile)
@app.route('/bookcount')
def bookcount():
    if not API_OK():
        return render_template('website_down.html')
    
    user_agent = request.headers.get('User-Agent')
    if any(keyword in user_agent.lower() for keyword in ['iphone', 'android', 'blackberry', 'mobile', 'linux']):
        if 'ubuntu' in user_agent.lower():
            mobile = ""
        else:
            mobile = "true"
    else:
        mobile = ""

    if request.args.get('languages') != None:
        language = request.args.get('languages')
    else:
        language = ""

    # Fetch JSON data from the endpoints
    url_support_languages = f'{api_url}/librarystats/v1/supported_languages/'
    response_support_languages = requests.get(url_support_languages)
    support_languages = response_support_languages.json()

    # For bookcount if language parameter is empty default to load:
    # Chinese, English, French, German, Japanese, Norwegian, Portugese, Russian, Spanish and Finnish
    if language == "":
        language = "zh,en,fr,de,ja,no,pt,ru,es,fi"

    url = f'{api_url}/librarystats/v1/bookcount/?languages={language}'
    response = requests.get(url)
    if response.status_code == 200:
        data = response.json()
    else:
        data = ""

    return stream_template('bookcount.html',
                            results=data,
                            language=language,
                            support_languages=support_languages,
                            mobile=mobile)

@app.route('/readership')
def readership():
    if not API_OK():
        return render_template('website_down.html')
    
    user_agent = request.headers.get('User-Agent')
    if any(keyword in user_agent.lower() for keyword in ['iphone', 'android', 'blackberry', 'mobile', 'linux']):
        if 'ubuntu' in user_agent.lower():
            mobile = ""
        else:
            mobile = "true"
    else:
        mobile = ""

    language = "en"

    url_support_languages = f'{api_url}/librarystats/v1/supported_languages/'
    response_support_languages = requests.get(url_support_languages)
    support_languages = response_support_languages.json()

    url = f'{api_url}/librarystats/v1/readership/{language}'
    response = requests.get(url)
    if response.status_code == 200:
        data = response.json()
    else:
        data = ""

    return stream_template('readership.html',
                            results=data,
                            language=language,
                            support_languages=support_languages,
                            mobile=mobile)

@app.route('/readership/<language>/')
def readership_language(language):
    if not API_OK():
        return render_template('website_down.html')
    
    user_agent = request.headers.get('User-Agent')
    if any(keyword in user_agent.lower() for keyword in ['iphone', 'android', 'blackberry', 'mobile', 'linux']):
        if 'ubuntu' in user_agent.lower():
            mobile = ""
        else:
            mobile = "true"
    else:
        mobile = ""

    if request.args.get('limit') != None:
        limit = request.args.get('limit')
    else:
        limit = ""

    url_support_languages = f'{api_url}/librarystats/v1/supported_languages/'
    response_support_languages = requests.get(url_support_languages)
    support_languages = response_support_languages.json()

    url = f'{api_url}/librarystats/v1/readership/{language}/?limit={limit}'
    response = requests.get(url)
    if response.status_code == 200:
        data = response.json()
    else:
        data = ""

    return stream_template('readership.html',
                            results=data,
                            language=language,
                            limit=limit,
                            support_languages=support_languages,
                            mobile=mobile)

@app.route('/status')
def status():
    if not API_OK():
        return render_template('website_down.html')
    
    response = requests.get(f'{api_url}/librarystats/v1/status/')
    status_data = response.json()
    return stream_template('status.html', status=status_data)

# special route to fetch and display the books from url on the web service
@app.route('/fetch_book')
def fetch_book_page():
    if not API_OK():
        return render_template('website_down.html')
    url = request.args.get('url')
    if not url:
        return jsonify({'error': 'Missing URL parameter'}), 400

    try:
        # Fetch HTML content of the book page
        response = requests.get(url)
        response.raise_for_status()
        soup = BeautifulSoup(response.content, 'html.parser')

        # Remove all images from the body
        for img in soup.body.find_all('img'):
            img.extract()

        # Store the body of the HTML with images removed
        body_content = str(soup.body)

        # Render the template with the fetched book.
        return render_template('book.html', body_content=body_content)
    except requests.RequestException as e:
        return jsonify({'error': str(e)}), 500

# Run Flask app
if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", port=os.getenv("PORT", default=5000))