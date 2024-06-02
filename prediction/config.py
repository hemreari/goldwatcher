import json

def load_config():
    try:
        config_file = open('../config.json', 'r')
    except:
        print("couldn't read the config file.")
        return

    try:
        config = json.load(config_file)
    except:
        print("couldn't load the config from the config file.")
        return
    return config