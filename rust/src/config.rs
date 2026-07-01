use anyhow::{Context, Result, anyhow};
use std::path::PathBuf;

use crate::opts::Opts;

#[derive(Debug)]
pub struct Config {
    pub pwd: PathBuf,
    pub config: PathBuf,
    pub operation: Operation,
}

impl TryFrom<Opts> for Config {
    type Error = anyhow::Error;

    fn try_from(value: Opts) -> Result<Self> {
        let operation = value.args.try_into()?;
        let pwd = get_pwd(value.pwd)?;
        let config = get_config(value.config)?;
        return Ok(Config {
            operation,
            config,
            pwd,
        });
    }
}

#[derive(Debug, PartialEq)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String),
    Remove(String),
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;
    fn try_from(value: Vec<String>) -> Result<Self> {
        let mut value = value;
        if value.len() == 0 {
            return Ok(Operation::Print(None));
        }

        let term = value.get(0).expect("expects to exist");
        if term == "add" {
            if value.len() != 3 {
                return Err(anyhow!(
                    "operation add expected 2 arguments ,got {}",
                    value.len() - 1
                ));
            }
            let mut arg = value.drain(1..=2);
            return Ok(Operation::Add(
                arg.next().expect("expect to exist"),
                arg.next().expect("expect to exist"),
            ));
        }
        if term == "rm" {
            if value.len() != 2 {
                return Err(anyhow!(
                    "operation remove expected 1 arguments ,got {}",
                    value.len() - 1
                ));
            }
            let arg = value.pop().expect("expect to exist");
            return Ok(Operation::Remove(arg));
        }
        if value.len() > 1 {
            return Err(anyhow!(
                "operation print expected 0 or 1 argument ,got {}",
                value.len()
            ));
        }
        let arg = value.pop().expect("expect to exist");
        return Ok(Operation::Print(Some(arg)));
    }
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = config {
        return Ok(v);
    }
    let mut loc = std::env::home_dir().context("unable to locate the home directory")?;

    loc.push("projector");
    loc.push("projector.json");

    return Ok(loc);
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = pwd {
        return Ok(v);
    }
    return Ok(std::env::current_dir().context("unable to lacate the current directory")?);
}

#[cfg(test)]
mod test {
    use crate::{
        config::{Config, Operation},
        opts::Opts,
    };
    use anyhow::Result;

    #[test]
    fn test_print_all() -> Result<()> {
        let cfg: Config = Opts {
            args: vec![],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(cfg.operation, Operation::Print(None));
        return Ok(());
    }
    #[test]
    fn test_print_key() -> Result<()> {
        let cfg: Config = Opts {
            args: vec!["foo".to_string()],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(cfg.operation, Operation::Print(Some("foo".to_string())));
        return Ok(());
    }

    #[test]
    fn test_add_key() -> Result<()> {
        let cfg: Config = Opts {
            args: vec![
                String::from("add"),
                String::from("foo"),
                String::from("bar"),
            ],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(
            cfg.operation,
            Operation::Add("foo".to_string(), "bar".to_string())
        );
        return Ok(());
    }
    #[test]
    fn test_remove_key() -> Result<()> {
        let cfg: Config = Opts {
            args: vec![String::from("rm"), String::from("foo")],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(cfg.operation, Operation::Remove("foo".to_string()));
        return Ok(());
    }
}
