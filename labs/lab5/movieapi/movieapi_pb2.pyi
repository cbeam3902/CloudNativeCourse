from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class MovieData(_message.Message):
    __slots__ = ["cast", "director", "title", "year"]
    CAST_FIELD_NUMBER: _ClassVar[int]
    DIRECTOR_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    cast: _containers.RepeatedScalarFieldContainer[str]
    director: str
    title: str
    year: int
    def __init__(self, title: _Optional[str] = ..., year: _Optional[int] = ..., director: _Optional[str] = ..., cast: _Optional[_Iterable[str]] = ...) -> None: ...

class MovieReply(_message.Message):
    __slots__ = ["cast", "director", "year"]
    CAST_FIELD_NUMBER: _ClassVar[int]
    DIRECTOR_FIELD_NUMBER: _ClassVar[int]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    cast: _containers.RepeatedScalarFieldContainer[str]
    director: str
    year: int
    def __init__(self, year: _Optional[int] = ..., director: _Optional[str] = ..., cast: _Optional[_Iterable[str]] = ...) -> None: ...

class MovieRequest(_message.Message):
    __slots__ = ["title"]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    title: str
    def __init__(self, title: _Optional[str] = ...) -> None: ...

class Status(_message.Message):
    __slots__ = ["code"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: int
    def __init__(self, code: _Optional[int] = ...) -> None: ...
