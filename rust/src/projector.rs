use serde::{Deserialize, Serialize};
use std::{collections::HashMap, path::PathBuf};

use crate::config::Config;

type HM = HashMap<String, String>;

#[derive(Debug, Serialize, Deserialize)]
pub struct Data {
    pub projector: HashMap<PathBuf, HM>,
}

pub struct Projector {
    pub config: Config,
    pub data: Data,
}

fn default_data() -> Data {
    return Data {
        projector: HashMap::new(),
    };
}

impl Projector {
    pub fn from_config(config: Config) -> Self {
        if std::fs::metadata(&config.config).is_ok() {
            let contents = std::fs::read_to_string(&config.config);
            let contents = contents.unwrap_or("{\"projector\":{}}".to_string());
            let data = serde_json::from_str(&contents);
            let data = data.unwrap_or(default_data());
            return Projector { config, data };
        }
        return Projector {
            config,
            data: default_data(),
        };
    }

    pub fn get_value(&self, key: String) -> Option<&String> {
        let mut curr = Some(self.config.pwd.as_path());
        let mut out = None;
        while let Some(p) = curr {
            if let Some(dir) = self.data.projector.get(p) {
                if let Some(value) = dir.get(&key) {
                    out = Some(value);
                }
            }
            curr = p.parent();
        }
        return out;
    }

    pub fn get_value_all(&self) -> Option<&HM> {
        let mut curr = Some(self.config.pwd.as_path());
        let mut paths = vec![];
        let mut out = None;
        while let Some(p) = curr {
            paths.push(curr);
            curr = p.parent();
        }
        for path in paths.into_iter().rev() {
            self.data.projector.get(path?).map(|x| out.insert(x));
        }
        return out;
    }

    pub fn set_value(&mut self, key: String, value: String) {
        self.data
            .projector
            .entry(self.config.pwd.clone())
            .or_default()
            .insert(key, value);
    }

    pub fn remove_valu(&mut self, key: String) {
        self.data
            .projector
            .entry(self.config.pwd.clone())
            .or_default()
            .remove(&key);
    }
}

#[cfg(test)]
mod test {}
